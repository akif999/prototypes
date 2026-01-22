from collections import Counter


def format_plain(items: list[tuple[str, int]]) -> str:
    return "\n".join(f"{name}: {count}" for name, count in items)


def format_table(items: list[tuple[str, int]]) -> str:
    if not items:
        return ""

    name_width = max(len(name) for name, _ in items)
    count_width = max(len(count) for count, _ in items)

    lines = []
    lines.append(f"{'Name':<{name_width}} {'Count':>{count_width}}")
    lines.append("-" * (name_width + count_width + 2))

    for name, count in items:
        lines.append(f"{name:<{name_width}} {count:>{count_width}}")

    return "\n".join(lines)
