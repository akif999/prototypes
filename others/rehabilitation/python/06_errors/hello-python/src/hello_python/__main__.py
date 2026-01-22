import argparse
import sys
from pathlib import Path

from hello_python.counter import count_names, InvalidCSVError


def main() -> None:
    parser = argparse.ArgumentParser(
        prog="hello-python",
        description="Count names from csv",
    )

    parser.add_argument(
        "--input",
        required=True,
        help="input csv file",
    )

    args = parser.parse_args()
    input_path = Path(args.input)

    try:
        counts = count_names(input_path)
    except FileNotFoundError as e:
        print(e, file=sys.stderr)
        sys.exit(1)
    except InvalidCSVError as e:
        print(f"Invalid CSV: {e}", file=sys.stderr)
        sys.exit(2)

    for name, count in counts.items():
        print(f"{name}: {count}")


if __name__ == "__main__":
    main()
