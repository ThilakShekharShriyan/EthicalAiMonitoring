# EthicalAiMonitoring


# Description
 write desc about what the project is 


## Features
What all is included in the projects
- gives overview of
- system design
- Program design paradigm
- Using DB
- AI ML using LLMs
- Distributed Systems

### Installations
- Clone the repo
- cp the env.example file make necessary changes
    ```
    cp common/config/.env.example common config/.env
    ```
- build and startservices make sure no containers/process are running on port (what ever we use)
        pkill -9 pid

# Project Structure

    Services
        - bias-detection : Service for detecting bias.
        - toxicity-analysis : Service for toxicity and harmful language detection.
        - api-gateway : Manages routing between services.


