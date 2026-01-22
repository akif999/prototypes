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

        names = [row["name"] for row in reader if row["name"]]

    return Counter(names)


def sort_counts(
    counts: Counter,
    *,
    by: str,
) -> list[tuple[str, int]]:
    if by == "name":
        return sorted(counts.items())
    elif by == "count":
        return sorted(counts.items(), key=lambda x: x[1], reverse=True)
    else:
        raise ValueError(f"Invalid sort key: {by}")
