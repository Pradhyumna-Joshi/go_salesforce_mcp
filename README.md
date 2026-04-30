# Salesforce MCP Server

A high-performance implementation of the Model Context Protocol (MCP) written in Go, designed to empower AI agents with deep, full-stack access to Salesforce environments. This server bridges the gap between natural language and Salesforce APIs, allowing for seamless data management, metadata exploration, Apex development, and Lightning UI orchestration.

-----

## 📑 Table of Contents

- [Overview](#overview)
- [Architecture](#️-architecture)
- [Key Modules & Tools](#-key-modules--tools)
  - [Org Management](#a-org-management-cli-bridge)
  - [Data & Search](#b-data--search-standard-rest)
  - [Metadata & Schema](#c-metadata--schema-introspection)
  - [Apex Development](#d-apex-development-tooling-api)
  - [Lightning UI](#e-lightning-ui-lwc--aura)
- [Core Design Decisions](#-core-design-decisions)
- [Getting Started](#️-getting-started)
- [Makefile Reference](#-makefile-reference)
- [Environment Configuration](#️-environment-configuration)
- [Security](#️-security)
- [Project Structure](#-project-structure)

-----

## Overview

The Salesforce MCP Server exposes the full Salesforce platform as a set of structured, AI-consumable tools via the Model Context Protocol. Instead of requiring an AI agent to know Salesforce API internals, this server abstracts authentication, schema introspection, SOQL/SOSL queries, Apex source management, and Lightning component editing into discrete, well-described tool calls.

**Key capabilities at a glance:**

- Zero-config authentication via Salesforce CLI session bridging
- Agent-centric prompting with embedded workflow instructions per tool
- Automatic JSON Schema generation via Go reflection — always in sync with code
- Context window optimization — raw Salesforce responses filtered to ~10% of original size
- Full-stack coverage: Org → Data → Metadata → Apex → LWC/Aura

-----

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────┐
│                IDFC Nexus                               │
└───────────────────────────┬─────────────────────────────┘
                            │ MCP (JSON-RPC over HTTP/SSE)
┌───────────────────────────▼─────────────────────────────┐
│                Salesforce MCP Server (Go)               │
│                                                         │
│  ┌─────────────┐  ┌──────────────┐  ┌────────────────┐  │
│  │ Org Mgmt    │  │Data/Metadata │  │ Apex / UI Dev  │  │
│  │(CLI Bridge) │  │ (REST API)   │  │ (Tooling API)  │  │
│  └──────┬──────┘  └──────┬───────┘  └───────┬────────┘  │
└─────────┼────────────────┼──────────────────┼───────────┘
          │                │                  │
┌─────────▼────────────────▼──────────────────▼───────────┐
│                    Salesforce Platform                  │
│          REST API   │   Tooling API   │   sf CLI        │
└─────────────────────────────────────────────────────────┘
```

|Component            |Detail                                                                 |
|---------------------|-----------------------------------------------------------------------|
|**Language**         |Go 1.21+                                                               |
|**Protocol**         |MCP (JSON-RPC over HTTP/SSE)                                           |
|**API Integration**  |Hybrid — Salesforce REST API, Tooling API, and `sf` CLI wrappers       |
|**Schema Generation**|Automated JSON Schema via Go reflection and `jsonschema` struct tags   |
|**Concurrency**      |Platform-specific signal handling (Unix/Windows) for graceful shutdowns|
|**Auth Strategy**    |Zero-config CLI session bridging — no credentials stored in-app        |

-----

## 🚀 Key Modules & Tools

### A. Org Management (CLI Bridge)

Integrates directly with the Salesforce CLI to manage full environment lifecycles. The server bridges `os/exec` calls to the `sf` binary, enabling the AI to authenticate and switch contexts without manual credential handling.

|Tool            |Description                                                                        |
|----------------|-----------------------------------------------------------------------------------|
|`list_orgs`     |Retrieves all Salesforce orgs currently authenticated in the local `sf` CLI        |
|`connect_to_org`|Initiates an OAuth web login flow (Production or Sandbox) and assigns a local alias|
|`open_org`      |Command-line trigger to open a specific org alias in the default web browser       |
|`disconnect_org`|Logs out and removes an org’s local session tokens                                 |

-----

### B. Data & Search (Standard REST)

Enables the AI to query and manipulate live Salesforce business data using the standard REST API. Supports both structured SOQL queries and cross-object SOSL keyword searches.

|Tool           |Description                                                                 |
|---------------|----------------------------------------------------------------------------|
|`query_sobject`|Executes standard SOQL queries to retrieve business records                 |
|`create_record`|Creates new records for any SObject (Account, Lead, Custom Object, etc.)    |
|`update_record`|Performs partial updates (PATCH) on existing records using their 18-digit ID|
|`global_search`|Uses SOSL to search for keywords across multiple objects simultaneously     |

-----

### C. Metadata & Schema (Introspection)

A suite of “verify-before-query” tools that give the AI a precise map of the org’s schema before attempting any data operation. Drastically reduces `INVALID_FIELD` and `INVALID_TYPE` errors caused by guessed field names.

|Tool                     |Description                                                                    |
|-------------------------|-------------------------------------------------------------------------------|
|`get_sobject_fields`     |Returns a filtered list of fields (Name, Label, Type) for a specific object    |
|`get_field_details`      |Drills into a specific field to retrieve picklist values and full metadata     |
|`get_child_relationships`|Maps parent-child object hierarchies — essential for complex JOIN-style queries|
|`get_record_types`       |Lists available Record Types, required for high-integrity record creation      |
|`get_system_context`     |Checks org limits (API usage, daily requests) to ensure safe execution         |


> **Design Note:** All describe responses are filtered server-side to return only `Name`, `Label`, and `Type` — reducing token cost by ~90% compared to a raw Salesforce Describe response.

-----

### D. Apex Development (Tooling API)

Transforms the AI into a full Salesforce developer. All operations are performed via the Tooling API, giving the agent the ability to read, write, execute, debug, and test Apex code end-to-end.

|Tool                    |Description                                                          |
|------------------------|---------------------------------------------------------------------|
|`apex_class_get`        |Retrieves the full source code of an existing Apex Class             |
|`apex_class_create`     |Creates a new Apex Class with provided source                        |
|`apex_class_update`     |Updates the source of an existing Apex Class                         |
|`apex_trigger_get`      |Retrieves the source and metadata of an Apex Trigger                 |
|`apex_trigger_create`   |Creates a new Apex Trigger bound to a specific SObject               |
|`apex_trigger_update`   |Updates the source of an existing Apex Trigger                       |
|`apex_execute_anonymous`|Executes ad-hoc Apex code blocks for testing and scripting           |
|`apex_logs_list`        |Lists available System Debug logs with timestamps and user context   |
|`apex_log_get_body`     |Fetches the raw body of a specific debug log for troubleshooting     |
|`apex_tests_run_async`  |Enqueues asynchronous unit test runs for specific classes or suites  |
|`apex_code_coverage_get`|Retrieves line-level coverage percentages for code quality validation|

-----

### E. Lightning UI (LWC & Aura)

Full-stack capabilities for both modern Lightning Web Components and legacy Aura components. Uses the shared Tooling API query engine from the Apex module, maintaining a DRY architecture.

|Tool                   |Description                                                                            |
|-----------------------|---------------------------------------------------------------------------------------|
|`lwc_bundle_list_files`|Lists all resources (JS, HTML, CSS, XML) inside an LWC bundle                          |
|`lwc_resource_get`     |Retrieves the source code of a specific file within an LWC bundle                      |
|`lwc_resource_update`  |Modifies a specific resource file within an LWC bundle                                 |
|`aura_definition_get`  |Fetches source code for Aura component definitions (Controller, Helper, Renderer, etc.)|

-----

## 🧠 Core Design Decisions

### I. Hybrid Authentication Strategy

Rather than implementing a custom OAuth flow, the server bridges the `sf` CLI via `os/exec`. AccessTokens are scraped from the CLI’s local credential store at runtime. This achieves **zero-config authentication** — if you’re logged in on your terminal, the AI agent is authenticated automatically.

### II. Reflection-Based Auto-Discovery

Go reflection and `invopop/jsonschema` struct tags are used to **auto-generate MCP tool definitions** at startup. This eliminates manual JSON schema maintenance and guarantees that the tool documentation exposed to the AI is always perfectly synchronized with the Go source code.

### III. Agent-Centric Prompt Engineering

Critical workflow instructions are **embedded directly in tool descriptions** rather than relying on external system prompts. Metadata tools are prefixed with `CRITICAL:` to enforce a verify-then-execute pattern — the AI checks the org’s schema before writing any query, preventing the common `INVALID_FIELD` and `__c` guessing errors.

### IV. Context Window Optimization

Raw Salesforce `Describe` API responses can exceed several megabytes per object, which would overflow an AI’s context window. All metadata handlers manually filter the response payload to return only essential data (`Name`, `Label`, `Type`), achieving a **~90% reduction in token cost** per introspection call.

### V. Cross-Module DRY Architecture

`ApexToolingQuery` is an exported, shared function used by both the Apex and Lightning UI modules. The LWC and Aura tools pass their SOQL strings to this shared Tooling API engine instead of reimplementing HTTP logic, maintaining clean separation of concerns across modules.

### VI. Platform-Agnostic Signal Handling

Separate Go build-tag files (`signal_unix.go`, `signal_windows.go`) handle OS-specific interrupt signals. This ensures that when the MCP server process is stopped, all open network sockets, `sf` CLI child processes, and HTTP connections are closed gracefully — preventing zombie processes and socket leaks across both Unix and Windows environments.

-----

## 🛠️ Getting Started

### Prerequisites

- **Salesforce CLI**: The `sf` executable must be installed and authenticated with at least one org.
- **Go**: Version 1.21 or higher.

### Installation

```bash
# Clone the repository
git clone https://github.com/Pradhyumna-Joshi/go_salesforce_mcp.git
cd go_salesforce_mcp

# Install dependencies
go mod tidy
```

### Quick Start

```bash
# Build and run in one step
make run
```

Or manually:

```bash
go run ./...
```

The server will start on the configured port and begin accepting MCP JSON-RPC connections over HTTP/SSE.

-----

## 📋 Makefile Reference

|Command     |Description                                      |
|------------|-------------------------------------------------|
|`make build`|Compiles the Go binary into the `./bin` directory|
|`make run`  |Builds and starts the MCP server                 |
|`make clean`|Removes all compiled build artifacts             |

-----

## ⚙️ Environment Configuration

Create a `.env` file in the root directory:

```env
SF_MCP_PORT=9000
```

|Variable     |Default|Description                       |
|-------------|-------|----------------------------------|
|`SF_MCP_PORT`|`9000` |The port the MCP server listens on|


> **Note:** The server does not store or manage Salesforce credentials directly. All authentication is delegated to the local `sf` CLI session store.

-----

## 🛡️ Security

|Concern                     |Approach                                                                                                                  |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------|
|**Credential Storage**      |Tokens managed exclusively through the `sf` CLI’s secure local store — never stored in-app                                |
|**Graceful Shutdown**       |OS interrupt signal traps ensure all sockets and child processes close cleanly on exit                                    |
|**Type Safety**             |Strict Go struct definitions with `jsonschema` tags prevent malformed payloads from reaching Salesforce APIs              |
|**Token Lifecycle**         |Session tokens are read at request time from the CLI store, respecting Salesforce’s native token expiry and refresh cycles|
|**No Credential Hardcoding**|Zero plaintext secrets — `.env` contains only port config, never API keys or passwords                                    |

-----

## 📂 Project Structure

```
.
├── main.go                     # Entry point & MCP server initialization
├── Makefile                    # Build, run, and clean commands
├── .env                        # Local environment configuration
│
├── tools/
│   ├── org/                    # Org Management tools (CLI bridge)
│   ├── data/                   # Data & Search tools (REST API)
│   ├── metadata/               # Metadata & Schema introspection tools
│   ├── apex/                   # Apex development tools (Tooling API)
│   └── ui/                     # Lightning UI tools (LWC & Aura)
│
├── client/
│   └── salesforce.go           # Core Salesforce HTTP client & auth bridge
│
├── signal_unix.go              # Unix-specific graceful shutdown (build tag)
└── signal_windows.go           # Windows-specific graceful shutdown (build tag)
```
