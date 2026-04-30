# Salesforce MCP Server

A high-performance implementation of the Model Context Protocol (MCP) written in Go, designed to empower AI agents with deep, full-stack access to Salesforce environments. This server bridges the gap between natural language and Salesforce APIs, allowing for seamless data management, metadata exploration, and development orchestration.

---

## 🏗️ Architecture

| Component | Detail |
|---|---|
| Language | Go (1.21+) |
| Protocol | MCP (JSON-RPC over HTTP/SSE) |
| API Integration | Hybrid — Salesforce REST API, Tooling API, and Salesforce CLI (`sf`) wrappers |
| Schema Generation | Automated JSON Schema via Go reflection and struct tags |
| Concurrency | Platform-specific signal handling (Unix/Windows) for graceful shutdowns |

---

## 🚀 Key Modules

### Org Management
Integrates directly with the Salesforce CLI to manage environment lifecycles.

| Tool | Description |
|---|---|
| `list_orgs` | Retrieves all Salesforce orgs authenticated in the local `sf` CLI |
| `connect_to_org` | Triggers browser-based OAuth flows for Production or Sandbox |
| `open_org` | Opens a specific org alias directly in the user's default browser |
| `disconnect_org` | Logs out and removes an org's local session tokens |

### Data & Search
Enables the AI to understand and manipulate the Salesforce database.

| Tool | Description |
|---|---|
| `query_sobject` | Executes standard SOQL queries to retrieve business records |
| `create_record` | Creates new records for any SObject (Account, Lead, Custom Object, etc.) |
| `update_record` | Performs partial updates (PATCH) on existing records using their 18-digit ID |
| `global_search` | Uses SOSL to search for keywords across multiple objects simultaneously |

### Metadata & Schema
"Verify-before-query" introspection tools that significantly reduce `INVALID_FIELD` errors.

| Tool | Description |
|---|---|
| `get_sobject_fields` | Returns a filtered list of fields (Name, Label, Type) for a specific object |
| `get_field_details` | Drills into a specific field to retrieve picklist values and metadata |
| `get_child_relationships` | Maps parent-child object hierarchies for complex queries |
| `get_record_types` | Lists available Record Types — required for high-integrity record creation |
| `get_system_context` | Checks org limits (API usage, daily requests) to ensure safe execution |

### Apex Development
Turns the AI into a full Salesforce developer.

| Tool | Description |
|---|---|
| `apex_class_get/create/update` | Full source code management for Apex Classes |
| `apex_trigger_get/create/update` | Source code and binding management for Apex Triggers |
| `apex_execute_anonymous` | Executes ad-hoc Apex code blocks for testing and scripting |
| `apex_logs_list` / `apex_log_get_body` | Fetches and reads raw System Debug logs |
| `apex_tests_run_async` | Enqueues unit test runs for classes or suites |
| `apex_code_coverage_get` | Retrieves line coverage percentages for code quality checks |

### Lightning UI (LWC & Aura)
Full-stack capabilities for modern and legacy UI frameworks.

| Tool | Description |
|---|---|
| `lwc_bundle_list_files` | Lists all resources (JS, HTML, CSS, XML) inside an LWC bundle |
| `lwc_resource_get/update` | Retrieves or modifies source code of specific files within an LWC |
| `aura_definition_get` | Fetches source code for Aura component definitions (Controller, Helper, etc.) |

---

## 🧠 Core Design Decisions

### I. Hybrid Authentication Strategy
The server bridges the `sf` CLI with the Go runtime via `os/exec`. By scraping the `AccessToken` from the CLI's local store, the server achieves **zero-config authentication** — if you're logged in on your terminal, the AI is logged in too.

### II. Reflection-Based Auto-Discovery
Go reflection and `jsonschema` struct tags are used to **auto-generate MCP tool definitions**. This keeps the AI's tool documentation perfectly synchronized with the underlying Go code at all times — no manual schema maintenance.

### III. Agent-Centric Prompt Engineering
Critical workflow instructions are **embedded directly in tool descriptions**. Metadata tools are prefixed with `CRITICAL` to force the AI to verify the schema before executing queries, preventing the common `INVALID_FIELD` errors that arise from guessing custom field names (`__c`).

### IV. Context Window Optimization
Raw Salesforce `Describe` API responses can be **megabytes in size**. All metadata handlers manually filter responses to return only essential fields (Name, Label, Type), reducing token cost by ~90% and preventing context window overflow.

### V. Cross-Module DRY Architecture
`ApexToolingQuery` is exported and shared across modules. The LWC and Aura tools reuse the Apex module's Tooling API query engine, eliminating duplicated HTTP logic and maintaining a clean separation of concerns.

### VI. Platform-Agnostic Signal Handling
Separate build-tag files handle Unix and Windows signal traps. This ensures all network sockets and local processes are closed gracefully when the server shuts down — no zombie processes or socket leaks.

---

## 🛠️ Getting Started

### Prerequisites

- **Salesforce CLI**: The `sf` executable must be installed and authenticated.
- **Go**: Version 1.21 or higher.

### Installation

Clone the repository and create a `.env` file in the root directory:

```env
SF_MCP_PORT=9000
