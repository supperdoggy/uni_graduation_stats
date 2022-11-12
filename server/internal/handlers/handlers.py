from fastapi import APIRouter, Depends, HTTPException


router = APIRouter()

@router.get("/")
def index():
    return "hi"
