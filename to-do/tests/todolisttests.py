import json
import os
import unittest
import sys
# sys.path.append(os.pardir)
import todolist


class TodoListTests(unittest.TestCase):
    def setUp(self) -> None:
        with open("./testdata.json", "r") as f:
            self.testdata_text = f.read()

    def tearDown(self) -> None:
        try:
            os.remove("./todo.json")
        except OSError:
            pass

    def create_data_file(self):
        with open("./todo.json", "w") as f:
            f.write(self.testdata_text)

    def create_todolist_and_safe_list(self):
        self.create_data_file()
        self.todolist = todolist.TodoList("./todo.json")
        self.list = json.loads(self.testdata_text)

    def get_data_file_as_string(self):
        with open("./todo.json", "r") as f:
            string = self.testdata_text = f.read()
        return string

    def test_save(self):
        self.create_todolist_and_safe_list()
        self.todolist.save()
        newfile_text = self.get_data_file_as_string()
        self.assertEqual(newfile_text, self.testdata_text)
