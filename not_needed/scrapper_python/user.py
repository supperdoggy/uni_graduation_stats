

import datetime
from typing import List


class WorkExperience:
    job_title: str
    duration: str
    company: str

    def __init__(self, job_title: str, duration: str, company: str) -> None:
        self.job_title = job_title
        self.duration = duration
        self.company = company

    def __str__(self):
        return f"{self.company}, {self.duration}, {self.job_title}"

class User:
    first_name: str
    last_name: str

    scraped: datetime.datetime
    experience: List[WorkExperience]

    def __init__(self, first_name: str, last_name: str, scraped: datetime.datetime, experience: List[WorkExperience]) -> None:
        self.first_name = first_name
        self.last_name = last_name
        self.scraped = scraped
        self.experience = experience

