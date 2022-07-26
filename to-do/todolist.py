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

    def id_less_string(self):
        string = ''
        for item in self.list:
            string = string + str(item.text) + '\n'
        string = string[:-1]
        return string

    def __len__(self):
        return len(self.list)

    def __iter__(self):
        return self.forward()

    def forward(self):
        current_item = 0
        while current_item < len(self):
            item = self.list[current_item]
            current_item += 1
            yield item

    def __getitem__(self, item_id):
        for todo in self.list:
            if todo.id == item_id:
                return todo
        raise IndexError("No such item")

    def __contains__(self, item):
        if item in self.list:
            return True
        else:
            return False

    def remove(self, item_id):
        for todo in self.list:
            if todo.id == item_id:
                self.list.remove(todo)
        raise IndexError("No such item")

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
                # print(item)
        except IOError:
            raise NoSuchListError

    def save(self):
        pass


class TodoItem:
    def __init__(self, id: int, text: str):
        self.id = id
        self.text = text

    def __str__(self):
        return self.text

    def asDict(self):
        return {'id': self.id, 'text': self.text}
