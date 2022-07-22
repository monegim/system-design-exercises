import json


class NoSuchListError(Exception): pass


class TodoList:
    def __init__(self, json_location):
        """
        Set up the list
        """
        self.json_location = json_location
        self.load()

    def add(self, text):
        try:
            last_id = self.list[-1].id
        except IndexError:
            last_id = -1

        new_item = TodoItem(last_id, text)
        self.list.append(new_item)
        return new_item

    def __str__(self):
        string_list = []
        for item in self.list:
            string_list.append(str(item.id) + '\t' + str(item.text))
        string = '\n'.join(string_list)
        return string

    def load(self):
        """
        Load the list data from file
        :return:
        """
        try:
            with open(self.json_location, 'r') as f:
                json_text = f.read()
            json_list = json.loads(json_text)
            self.list = []
            for item in json_list:
                self.list.append(TodoItem(item['id'], item['text']))
                print(item)
        except IOError:
            raise NoSuchListError


class TodoItem:
    def __init__(self, id: int, text: str):
        self.id = id
        self.text = text

    def __str__(self):
        return self.text

    def asDict(self):
        return {'id': self.id, 'text': self.text}
