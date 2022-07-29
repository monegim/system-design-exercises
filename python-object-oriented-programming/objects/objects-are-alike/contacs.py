from typing import List


class Contact:
    all_contacts: List['Contact'] = []

    def __init__(self, name:str, email: str) -> None:
        self.name = name
        self.email = email
        Contact.all_contacts.append(self)

    def __repr__(self) -> str:
        return (
            f"{self.__class__.__name__}("
            f"{self.name!r}, {self.email!r}"
            f")"
        )

class Supplier(Contact):
    def order(self, order: "Order") -> None:
        print(
            "If this were a real system we could send "
            f"'{order}' order to '{self.name}'"
        )

c_1 = Contact("Dusty", "dusty@example.com")
c_2 = Contact("Steve", "steve@itmaybeahack.com")

class Friend(Contact):
    def __init__(self, name: str, email: str, phone: str) -> None:
        super().__init__(name, email)
        self.phone = phone