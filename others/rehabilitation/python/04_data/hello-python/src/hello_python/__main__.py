import argparse
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(
        prog="hello-python",
        description="greeting from file",
    )

    parser.add_argument(
        "--input",
        required=True,
        help="input text file",
    )
    parser.add_argument(
        "--output",
        required=True,
        help="output text file",
    )

    args = parser.parse_args()

    input_path = Path(args.input)
    output_path = Path(args.output)

    lines = input_path.read_text(encoding="utf-8").splitlines()

    greetings = [f"Hello, {line}" for line in lines]

    output_path.write_text(
        "\n".join(greetings),
        encoding="utf-8"
    )


if __name__ == "__main__":
    main()
