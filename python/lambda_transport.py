import ctypes


def wrap_function(lib, funcname, restype, argtypes):
    """Simplify wrapping ctypes functions"""
    func = getattr(lib, funcname)
    func.restype = restype
    func.argtypes = argtypes
    return func


gorpc = ctypes.cdll.LoadLibrary('gorpc.so')


Handler = ctypes.CFUNCTYPE(
    ctypes.c_void_p,
    ctypes.POINTER(ctypes.c_wchar_p),
    ctypes.POINTER(ctypes.c_wchar_p))

Start = wrap_function(gorpc, 'Start', None, [Handler])


def start(handler):
    Start(Handler(handler))
