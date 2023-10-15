import logging
from in_bound.grpc.server import serve

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    serve("50050")

