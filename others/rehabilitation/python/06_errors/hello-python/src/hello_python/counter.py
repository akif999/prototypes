import csv
from collections import Counter
from pathlib import Path


def count_names(csv_path: Path) -> Counter:
    with csv_path.open(encoding="utf-8", newline="") as f:
        reader = csv.DictReader(f)
        names = [row["name"] for row in reader]

    return Counter(names)
