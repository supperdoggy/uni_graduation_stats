from typing import Union

from fastapi import FastAPI
from internal.handlers import handlers

app = FastAPI()
app.include_router(handlers.router)


