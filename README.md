# EthicalAiMonitoring


# Description

## What is Ethical AI
Ethical AI is artificial intelligence that adheres to well-defined ethical guidelines regarding fundamental values, including such things as individual rights, privacy, non-discrimination, and non-manipulation

## Why Ethical AI?
AI has the potential to be used for both good and evil purposes. The benefits from the ethical uses of AI are numerous and significant. The application of AI can help organizations operate more efficiently, produce cleaner products, reduce harmful environmental impacts, increase public safety, and improve human health.

We need to make sure we use them responsible, having some Laws and regulations around them so everyone follows guidelines and can adhere to them.

## Key Considerations
- Fairness: AI systems should not unfairly discriminate against individuals or groups. 
- Transparency: AI systems should be transparent and explainable so that users and developers can understand them
- Accountability: AI systems should be auditable and traceable.
- Privacy: AI systems should protect personal data
- Security: AI systems should be protected from unwanted harms and vulnerabilities to attack
- Human oversight: Humans should retain ultimate responsibility and accountability for AI systems



# Features
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


# Todo
- Read about Ethical AI monitoring and exisiting solutions
- Come up with modules and services to add
- System design the layout
- Program design and  

# Janavi
- Detecting Hallucinations 
- Setup ollama Service


# Thilak
- Detecting Toxicity
- PVC Database.

