#!/bin/python

from jinja2 import Template

COPY_INTO = '''
copy into {{table_name}}
  from @my_json_stage/{{table_name}}
  file_format = (format_name = my_json_format)
  on_error = 'skip_file';
'''

def main():
  print("Hello World!")

if __name__ == "__main__":
  main()