
from collections import OrderedDict


class Car:
    def __init__(self, registration_number, color):
        self.registration_number = registration_number
        self.color = color

    def __str__(self):
        return "Car [regNo=" + self.registration_number + ", color=" + self.color + "]"


class Parking_lot(Car):
    # we need to maintain the orders of cars while showing 'status'
    slotTocar = OrderedDict()
    i = 0

    def __init__(self, total_slots):
        # intialize all slots as free
        self.free_slots = [i for i in range(1, total_slots + 1)]
        # zero occupied initally
        self.occupied_slots = []

    def park_new_car(self, registration_number, color):
        self.slotTocar[self.next_free_slot()] = Car(registration_number, color)

    def next_free_slot(self):
        self.i += 1
        return self.i

    def status(self):
        print("SLOT",       "REG_NUM",        "COLOR")
        for slot, car in self.slotTocar.items():
            print(slot,       car.registration_number,        car.color)


if __name__ == "__main__":
    p = Parking_lot(10)
    p.park_new_car("ABC", "white")
    p.park_new_car("CDE", "Black")
    p.park_new_car("ABXYZC", "RED")
    p.status()
