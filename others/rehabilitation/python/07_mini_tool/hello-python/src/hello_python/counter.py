import csv
from collections import Counter
from pathlib import Path


class InvalidCSVError(Exception):
    pass


def count_names(csv_path: Path) -> Counter:
    if not csv_path.exists():
        raise FileNotFoundError(f"File not found: {csv_path}")

    with csv_path.open(encoding="utf-8", newline="") as f:
        reader = csv.DictReader(f)

        if "name" not in reader.fieldnames:
            raise InvalidCSVError("CSV must contain 'name' column")

        names = []
        for row in reader:
            if not row["name"]:
                continue
            names.append(row["name"])

    return Counter(names)
