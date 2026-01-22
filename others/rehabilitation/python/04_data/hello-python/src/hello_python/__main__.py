import argparse
import csv
from collections import Counter
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(
        prog="hello-python",
        description="Count names from csv",
    )

    parser.add_argument(
        "--input",
        required=True,
        help="input csv file",
    )
    # parser.add_argument(
    #     "--output",
    #     required=True,
    #     help="output text file",
    # )

    args = parser.parse_args()
    input_path = Path(args.input)

    with input_path.open(encoding="utf-8", newline="") as f:
        reader = csv.DictReader(f)
        names = [row["name"] for row in reader]

    counts = Counter(names)

    for name, count in counts.items():
        print(f"{name}: {count}")


if __name__ == "__main__":
    main()
