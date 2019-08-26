from redis_queue import doSquare
from rq import Queue
from redis import Redis


squareQ = Queue(connection=Redis())

if __name__ == "__main__":
    arr = [1, 2, 3, 4, 5, 6]
    for v in arr:
        squareQ.enqueue(doSquare, v)
