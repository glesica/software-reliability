import ast
from typing import Any, Iterable, TextIO

    
def _is_camel(value: str) -> bool:
    if "_" in value:
        return False
    if value[0].upper() != value[0]:
        return False
    
    return True


def _is_snake(value: str) -> bool:
    if value.lower() != value:
        return False
    
    return True


class CaseVisitor(ast.NodeVisitor):
    def __init__(self):
        self.errors = []

    def visit_ClassDef(self, node: ast.ClassDef) -> Any:
        if not _is_camel(node.name):
            self.errors.append(f"class {node.name} should be CamelCase")
        return super().generic_visit(node)

    def visit_FunctionDef(self, node: ast.FunctionDef) -> Any:
        if not _is_snake(node.name):
            self.errors.append(f"function {node.name} should be snake_case")
        return super().generic_visit(node)


def check_file(file: TextIO) -> Iterable[str]:
    code = file.read()
    tree = ast.parse(code)

    case_visitor = CaseVisitor()
    case_visitor.visit(tree)

    return case_visitor.errors


if __name__ == "__main__":
    import sys

    path = sys.argv[1]
    with open(path, "r") as file:
        errors = check_file(file)
    
    exit_code = 0
    for error in errors:
        exit_code = 1
        print(f"{path}: {error}")
    
    sys.exit(exit_code)
