from typing import Optional


class LongNameDict(dict[str, int]):
    def longest_key(self)-> Optional[str]:
        longest = None
        for key in self:
            if longest is None or len(key) > len(longest):
                longest = key
        return longest

articles_read = LongNameDict()
articles_read['mostafa'] = 5
articles_read['dariush'] = 7
articles_read['kianoosh_a'] = 1

longest = articles_read.longest_key()
print(longest)
print(max(articles_read, key=len))

        