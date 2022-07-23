import os
import unittest


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

