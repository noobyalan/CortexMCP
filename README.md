# CortexMCP 

## I. Introduction

The rapid development of artificial intelligence has made technologies like Large Language Models (LLMs), Agents, and Model Context Protocols (MCPs) critical drivers of AI application development. These technologies collaborate to provide robust support for building intelligent and efficient AI applications.


## II. Core Concepts

### A. Large Language Model (LLM)

LLMs are natural language processing tools based on deep learning techniques, trained on extensive datasets to provide strong language understanding and generation capabilities. For example, when a user inquires about the weather, the LLM can interpret the query's semantics and generate a relevant response. Common LLMs include the GPT series and others like Wenxin Yiyan and Deepseek.
<img width="1460" height="458" alt="image" src="https://github.com/user-attachments/assets/b397789b-18b8-4e51-b430-3cff44806ed6" />
### B. Agent

An Agent is an AI system that can autonomously perceive its environment, make decisions, and take actions to achieve specific goals. On top of a LLM, an Agent utilizes tools or plugins to enhance its interaction with the real world. For instance, when a user requests external information, the Agent interprets the question using the LLM and calls relevant tools to accomplish the task, such as fetching inter-store communication data.
<img width="1734" height="934" alt="image" src="https://github.com/user-attachments/assets/cf0c63b8-cd55-4072-939b-ea081e21a468" />


### C. Model Context Protocol (MCP)

MCP is an open standard that connects AI assistants with external systems, facilitating standardized communication between AI models and external data sources or tools. MCP servers provide a unified interface, simplifying and enhancing the integration of AI with various systems. In development, MCP can help retrieve contextual information, such as user interaction history, to better understand intent and provide more accurate responses.
![Uploading image.png…]()


## III. Development Process

### A. Environment Setup

1. **Select the Technology Stack**: Prepare your development environment with required libraries and frameworks for LLMs, Agents, and MCPs.
  
### B. Build the AI Application

1. **Define Objectives**: Clarify the application scope and required functionalities.
2. **Choose and Configure the LLM**: Select a suitable model based on your needs, configuring it for interaction.
3. **Design the Agent**: Outline how the Agent operates, including decision-making and response processes.

### C. Integrate MCP

1. **Build the MCP Server**: Establish an MCP server to manage communication and tool registration.
2. **Register Tools**: Insert the tools that the Agent will use for interaction.

### D. Implement Business Logic

Develop the necessary code to handle user requests, process inputs, and call tools to generate appropriate responses using the LLM.

### E. Testing and Optimization

Conduct thorough testing to ensure that all components interact correctly and efficiently, and apply optimizations based on performance outcomes.

## IV. Integrating LLMs

### A. Ark Model Integration

1. Access the Ark console to apply for model usage.
2. Retrieve the API Key and endpoint information necessary for API calls.

### B. Ollama Model Integration

1. Download and install Ollama.
2. Pull the required model and ensure proper local setup.

## V. Building the Agent

### A. Design and Create the Agent

1. Establish how the Agent will interact with users and manage tool calls.
2. Ensure the Agent can register and utilize tools provided by the MCP effectively.

## VI. Invoking the Agent

When users send queries, leverage the Agent to analyze the request, interact with the relevant tools through MCP, and return the necessary information.

## VII. Appendix

### A. RAG (Retrieval-Augmented Generation)

RAG combines information retrieval with text generation, effectively enhancing response accuracy and information richness. Within an Agent setup, RAG acts as a knowledge component, enhancing decision-making capabilities by providing reliable information.


## Start the project

macOS/Linux： `APP_ID=<app_id> APP_SECRET=<app_secret> ./bootstrap.sh`

Windows： `set APP_ID=<app_id>&set APP_SECRET=<app_secret>&bootstrap.bat`
