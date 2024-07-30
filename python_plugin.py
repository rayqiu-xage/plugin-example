from jsonrpcserver import method, serve

class GreeterPlugin:
    @method
    def greet(self, name):
        return f"Hello {name} from the Python plugin GreeterPlugin."

serve(GreeterPlugin(), address="localhost", port=8080)
