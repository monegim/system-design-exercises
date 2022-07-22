from enum import Enum
from abc import ABCMeta, abstractmethod


class Rank(Enum):
    OPERATOR = 0
    SUPERVISOR = 1
    DIRECTOR = 2


class Employee(metaclass=ABCMeta):
    def __init__(self, employee_id, name, rank, call_center):
        self.employee_id = employee_id
        self.name = name
        self.rank = rank
        self.call = None
        self.call_center = call_center

    def take_call(self, call):
        """Assume"""
        self.call = call
        self.call.employee = self
        self.call.state = CallState.IN_PROGRESS

    def complete_ccall(self):
        self.call.state = CallState.COMPLETE
        self.call_center.notify_call_completed(self.call)

    @abstractmethod
    def escalate_call(self):
        pass

    def _escalate_call(self):
        self.call.state = CallState.READY
        call = self.call
        self.call = None
        self.call_center.notify_call_completed(call)


class Operator(Employee):
    def __init__(self, employee_id, name):
        super(Operator, self).__init__(employee_id, name, Rank.OPERATOR)

    def escalate_call(self):
        self.call.level = Rank.DIRECTOR
        self._escalate_call()


class Supervisor(Employee):

    def __init__(self, employee_id, name):
        super(Supervisor, self).__init__(employee_id, name, Rank.SUPERVISOR)

    def escalate_call(self):
        self.call.level = Rank.DIRECTOR
        self._escalate_call()


class Director(Employee):

    def __init__(self, employee_id, name):
        super(Director, self).__init__(employee_id, name, Rank.DIRECTOR)

    def escalate_call(self):
        raise NotImplementedError('Directors must be able to handle any call')


class CallState(Enum):
    READY = 0
    IN_PROGRESS = 1
    COMPLETE = 2


class Call(object):
    def __init__(self, rank):
        self.state = CallState.READY
        self.rank = rank
        self.employee = None

class CallCenter:
    def __init__(self, operators):
        self.operators = operators