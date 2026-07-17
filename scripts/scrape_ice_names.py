#!/usr/bin/env python3
"""
Scrapes ICE train names from elektrolok.de and inserts them into the
trainspotter database.

Only the currently active name (no end date set) per TZ is used.
If a TZ already exists in the DB it is skipped.

Requirements:
    pip install requests beautifulsoup4 psycopg2-binary

Usage:
    DATABASE_URL=postgresql://trainspotter:trainspotter@localhost:5432/trainspotter \
        python3 scripts/scrape_ice_names.py
"""

import os
import re
import sys

import psycopg2
import requests
from bs4 import BeautifulSoup

URL = "https://www.elektrolok.de/statistiken/taufe.php?art=ice"


def scrape() -> list[dict]:
    response = requests.get(URL, timeout=30)
    response.raise_for_status()

    soup = BeautifulSoup(response.text, "html.parser")

    entries = []
    current_baureihe = None

    # The main data table starts after the "ICE-Triebwagen" header
    main_table = None
    for table in soup.find_all("table"):
        if table.find(string=re.compile("ICE-Triebwagen")):
            main_table = table
            break

    if main_table is None:
        print("Could not find main data table", file=sys.stderr)
        return []

    for row in main_table.find_all("tr"):
        cells = row.find_all("td")

        # Baureihe section header (single cell spanning multiple columns)
        if len(cells) == 1 or (len(cells) > 0 and cells[0].get("colspan")):
            text = cells[0].get_text(strip=True)
            m = re.search(r"Baureihe\s+(\S+)", text)
            if m:
                current_baureihe = m.group(1)
            continue

        # Data row: Lok | Name | vom | bis | Bemerkung
        if len(cells) < 4:
            continue

        lok = cells[0].get_text(strip=True)
        name = cells[1].get_text(strip=True)
        bis = cells[3].get_text(strip=True)

        # Skip header row and empty names
        if not lok or not name or name == "Name" or lok == "Lok":
            continue

        # Only keep currently active names (no end date)
        if bis:
            continue

        tz_match = re.search(r"Tz\s+(\d+)", lok)
        if not tz_match:
            continue

        entries.append(
            {
                "tz": int(tz_match.group(1)),
                "baureihe": current_baureihe or "",
                "name": name,
            }
        )

    return entries


def insert(entries: list[dict]) -> None:
    database_url = os.environ.get("DATABASE_URL")
    if not database_url:
        print("DATABASE_URL environment variable is not set", file=sys.stderr)
        sys.exit(1)

    conn = psycopg2.connect(database_url)
    cur = conn.cursor()

    inserted = 0
    skipped = 0

    for entry in entries:
        cur.execute("SELECT 1 FROM trains WHERE tz = %s", (entry["tz"],))
        if cur.fetchone():
            skipped += 1
            continue

        cur.execute(
            "INSERT INTO trains (tz, baureihe, name) VALUES (%s, %s, %s)",
            (entry["tz"], entry["baureihe"], entry["name"]),
        )
        inserted += 1

    conn.commit()
    cur.close()
    conn.close()

    print(f"Done: {inserted} inserted, {skipped} skipped (already existed)")


if __name__ == "__main__":
    print(f"Fetching data from {URL} ...")
    entries = scrape()
    print(f"Found {len(entries)} active ICE entries")

    if not entries:
        sys.exit(0)

    insert(entries)
