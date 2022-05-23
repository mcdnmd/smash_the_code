import os
from os.path import join
from typing import Tuple, List

ROOT_DIR_PATH = os.path.dirname(__file__)
ROOT_DIR = ROOT_DIR_PATH.split('/')[-1]


OUTPUT_FILE = 'codingame.txt'
EXCEPT_DIRS = ['test', 'venv']
EXCEPT_FILES = ['monofile.go', OUTPUT_FILE]


def find_all_files(path: str, files: list):
    for file in os.listdir(path):
        if os.path.isdir(file) and file not in EXCEPT_DIRS and not file.startswith('.'):
            find_all_files(join(path, file), files)
        elif file.endswith('.go') and file not in EXCEPT_FILES:
            files.append(join(path, file))


def find_import_text(text: str):
    start = text.find('import')
    end = text.find(')')
    end_by_type = text.find('type')
    end_by_func = text.find('func')
    if end_by_func > 0 or end_by_func > 0:
        selection = [i for i in [end_by_type, end_by_func, end] if i > 0]
        end = min(selection)
    data = text[start:end]
    return data


def parse_import(filepath: str) -> Tuple[List[str], List[str]]:
    with open(filepath, 'r') as f:
        text = f.read()
        data = find_import_text(text)
        data = data.replace('import', '')
        data = data.replace('(', '')
        data = data.replace(')', '')
        data = data.replace('\n', '')
        data = data.replace('\t', '')
        data = data.replace('"', ' ')
        data = data.strip()
        print(data)
        used_libs = data.split()
        project_import = [i.split('/')[-1] for i in used_libs if i.startswith(ROOT_DIR)]
        other_import = [i for i in used_libs if not i.startswith(ROOT_DIR)]
    return project_import, other_import


def write_line_in_output(s: str, output, p_libs):
    if s.strip().startswith('//') or\
            s.startswith('package') or\
            s.startswith('import') or \
            (s.strip().startswith('"') and s.strip().endswith('"') and s.count('"') == 2) or\
            (s.startswith(')')):
        return
    for lib in p_libs:
        if s.count(lib) > 0:
            lib = f'{lib}.'
            s = s.replace(lib, '')
    output.write(s)


def work_with_file(filepath: str, output, project_import):
    with open(filepath) as f:
        line = f.readline()
        while line != '':
            write_line_in_output(line, output, project_import)
            line = f.readline()


def insert_file_header(import_libs: set):
    with open(OUTPUT_FILE) as input_file:
        content = input_file.readlines()

    content.insert(0, 'package main\n\n')
    libs = "\n".join([f'\t"{lib}"' for lib in sorted(import_libs)])
    content.insert(1, f'import (\n{libs}\n)\n')

    with open(OUTPUT_FILE, 'w') as output_file:
        output_file.writelines(content)


if __name__ == '__main__':
    files = []
    find_all_files(ROOT_DIR_PATH, files)
    import_libs = set()

    with open(OUTPUT_FILE, 'w') as output:
        for filepath in files:
            print(filepath)
            project_import, other_import = parse_import(filepath)
            for lib in other_import:
                import_libs.add(lib)
            work_with_file(filepath, output, project_import)
    insert_file_header(import_libs)
