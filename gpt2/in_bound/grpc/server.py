import logging
import sys
import grpc
from concurrent import futures

sys.path.append("./")
from in_bound.grpc.proto.stubs import gpt2_pb2, gpt2_pb2_grpc

from transformers import pipeline, set_seed


class GenerateText(gpt2_pb2_grpc.ChatServiceServicer):
    def GenerateText(self, request, context):
        logging.info("request: {" + str(request) + "}")
        generator = pipeline('text-generation', model='gpt2')
        set_seed(42)
        response = generator(request.text, max_length=30, num_return_sequences=5)
        response = list(map(lambda x: str(x['generated_text']), response))
        logging.info("response: " + str(response) + "}")
        return gpt2_pb2.TextResponse(text=response)


def serve(port, max_workers=10):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=max_workers))
    gpt2_pb2_grpc.add_ChatServiceServicer_to_server(GenerateText(), server)
    server.add_insecure_port("[::]:" + port)

    server.start()
    logging.info("Server started, listening on " + port)
    server.wait_for_termination()
