from __future__ import annotations
from typing import List

class ContactList(List["Contact"]):
    def search(self, name: str) -> list["Contact"]:
        matching_contacts: List["Contact"] = []
        for contact in self:
            if name == contact.name:
                matching_contacts.append(contact)
        return matching_contacts

class Contact:
    pass