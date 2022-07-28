from __future__ import annotations
from operator import concat
from typing import List

class ContactList(List["Contact"]):
    def search(self, name: str) -> list["Contact"]:
        matching_contacts: List["Contact"] = []
        for contact in self:
            if name == contact.name:
                matching_contacts.append(contact)
        return matching_contacts

class Contact:
    all_contacts = ContactList()

    def __init__(self, name: str, email: str) -> None:
        self.name = name
        self.email = email
        Contact.all_contacts.append(self)
    def __repr__(self) -> str:
        return (
        f"{self.__class__.__name__}("
        f"{self.name!r}, {self.email!r}" f")"
)