# Scalable Upload Orchestration Service

A high-performance, distributed backend service built in Go to orchestrate the parallel, resumable upload of terabyte-scale files to cloud storage. This project is a deep dive into building a cloud-native, fault-tolerant system using modern system design principles.

---
## Key Features

* **Parallel Chunking:** Splits large files into chunks for high-throughput parallel uploads.
* **Stateless & Scalable:** Designed to be horizontally scalable with an externalized state store.
* **Fault-Tolerant & Resumable:** Recovers from network interruptions and failures seamlessly.
* **Secure Orchestration:** Uses S3 presigned URLs to offload data transfer, minimizing server load and enhancing security.

---
## Tech Stack

* **Backend:** Go (Golang), Gin
* **State Store:** Redis
* **Cloud & Deployment:** AWS S3, Docker

---
## Getting Started

### Prerequisites

* Go 1.18+
* Docker & Docker Compose
* AWS Account and configured credentials

### Installation & Setup

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/alpha-abhii/parallel-uploader.git](https://github.com/alpha-abhii/parallel-uploader.git)
    cd parallel-uploader
    ```

2.  **Configure Environment:**
    * Create a `configs/config.yaml` file based on a provided template (we will create this later).
    * Set up your AWS credentials locally.

3.  **Run the application:**
    ```bash
    docker-compose up --build
    ```

4.  **The service will be available at:**
    `http://localhost:8080`