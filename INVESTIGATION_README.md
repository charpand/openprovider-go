# Domain Transfer Investigation Results

This directory contains the complete investigation for implementing domain transfer support in the Terraform Openprovider provider.

## 📋 Investigation Documents

### 1. [TRANSFER_SUMMARY.md](TRANSFER_SUMMARY.md) - **START HERE**
Quick reference guide with key findings and decisions.
- **Best for:** Quick overview, decision summary, next steps
- **Read time:** 5 minutes

### 2. [DOMAIN_TRANSFER_INVESTIGATION.md](DOMAIN_TRANSFER_INVESTIGATION.md) - **DETAILED ANALYSIS**
Comprehensive investigation document with full analysis.
- **Best for:** Understanding the complete analysis, design rationale, technical details
- **Read time:** 20-30 minutes
- **Contents:**
  - Executive summary
  - Current state analysis
  - Openprovider API analysis
  - Implementation approaches (with pros/cons)
  - Transfer workflow design
  - Technical requirements
  - Configuration examples
  - Scope definition (in/out)
  - Risk analysis and mitigations
  - Implementation phases
  - Success criteria

### 3. [DRAFT_ISSUE_DOMAIN_TRANSFER.md](DRAFT_ISSUE_DOMAIN_TRANSFER.md) - **IMPLEMENTATION PLAN**
Ready-to-use GitHub issue for implementation.
- **Best for:** Creating implementation ticket, development roadmap
- **Read time:** 15-20 minutes
- **Contents:**
  - Detailed requirements
  - API client implementation specs
  - Provider resource implementation specs
  - Testing strategy
  - Documentation requirements
  - Phase-by-phase implementation plan
  - Code examples and schemas
  - Acceptance criteria per phase
  - Success metrics

## 🎯 Key Findings

### ✅ Recommendation: FEASIBLE AND RECOMMENDED

**Approach:** Create a new `openprovider_domain_transfer` resource (separate from existing `openprovider_domain`)

**API Support:** Fully supported via `POST /v1beta/domains/transfer` endpoint

**Effort:** 15-20 hours for complete implementation

### 🔑 Key Decision

**Separate Resource vs. Extending Existing Resource**

We recommend creating a **separate `openprovider_domain_transfer` resource** because:
- ✅ Clear intent and semantics (transfer vs. register)
- ✅ Easier to handle transfer-specific parameters
- ✅ No risk to existing domain resource
- ✅ Follows Terraform best practices
- ✅ Better error handling for transfer scenarios

## 🚀 Proposed Resource

```hcl
resource "openprovider_domain_transfer" "example" {
  # Required
  domain       = "example.com"
  auth_code    = var.auth_code  # EPP/authorization code from current registrar
  owner_handle = openprovider_customer.owner.handle
  
  # Optional
  admin_handle   = openprovider_customer.admin.handle
  tech_handle    = openprovider_customer.tech.handle
  billing_handle = openprovider_customer.billing.handle
  autorenew      = true
  ns_group       = "my-ns-group"
  
  # Computed (read-only)
  id              = "123456"     # Domain ID in Openprovider
  status          = "REQ"        # Transfer status (REQ → ACT)
  expiration_date = "2025-12-31" # Expiration after transfer
}
```

## 📊 Scope Summary

### ✅ In Scope
- Domain transfer initiation with auth code
- Contact handle configuration
- Nameserver configuration (ns_group or nameservers)
- Transfer status tracking
- Import existing transferred domains
- WHOIS privacy configuration
- Post-transfer domain management

### ❌ Out of Scope (Initial Release)
- Transfer approval (when Openprovider is losing registrar)
- Auth code retrieval for outbound transfers
- Bulk transfer operations
- Active status polling
- Advanced TLD-specific fields

## 📈 Implementation Phases

1. **Phase 1: API Client (MVP)** - Core transfer API functions
2. **Phase 2: Provider Resource (Core)** - Basic resource implementation
3. **Phase 3: Full Features** - Update, import, validation, all options
4. **Phase 4: Documentation** - Templates, examples, guides
5. **Phase 5: Testing & QA** - Comprehensive test coverage

**Total Estimate:** 15-20 hours

## 🔒 Security Considerations

- Auth codes are sensitive and should be in variables
- Auth codes appear in Terraform state file
- Recommend encrypted backend for production
- Mark auth_code as sensitive in schema

## 🎬 Transfer Workflow

```
┌─────────────────────────────────────┐
│ User obtains auth code from current│
│ registrar                           │
└──────────┬──────────────────────────┘
           │
           ▼
┌─────────────────────────────────────┐
│ User creates Terraform config       │
│ with openprovider_domain_transfer   │
└──────────┬──────────────────────────┘
           │
           ▼
┌─────────────────────────────────────┐
│ terraform apply                     │
│ POST /v1beta/domains/transfer       │
└──────────┬──────────────────────────┘
           │
           ▼
┌─────────────────────────────────────┐
│ Openprovider initiates transfer     │
│ Status: REQ (Requested)             │
│ Timeline: 5-7 days typical          │
└──────────┬──────────────────────────┘
           │
           ▼
┌─────────────────────────────────────┐
│ User: terraform refresh             │
│ Check status attribute              │
└──────────┬──────────────────────────┘
           │
           ▼
┌─────────────────────────────────────┐
│ Transfer complete                   │
│ Status: ACT (Active)                │
│ Domain at Openprovider              │
└─────────────────────────────────────┘
```

## 📝 Important Notes

1. **Async Operation:** Transfers take 5-7 days; resource creation succeeds when initiated, not completed
2. **Status Tracking:** Monitor via `status` attribute (REQ → ACT)
3. **Delete Behavior:** `terraform destroy` removes from state only, does NOT delete the domain
4. **No Breaking Changes:** Existing resources are unaffected
5. **Import Support:** Existing transferred domains can be imported

## ✅ Success Criteria

- [x] Clear, comprehensive investigation document
- [x] API analysis complete
- [x] Implementation approach defined
- [x] Workflow documented
- [x] Draft issue ready for implementation
- [x] No technical blockers identified
- [x] Risk level assessed as LOW

## 🎯 Next Steps

1. **Review** these investigation documents with stakeholders
2. **Approve** the recommended approach (separate resource)
3. **Create GitHub issue** from `DRAFT_ISSUE_DOMAIN_TRANSFER.md`
4. **Begin implementation** following the phased plan

## 📚 Additional References

- [Openprovider API Documentation](https://docs.openprovider.com/swagger.json)
- [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework)
- [ICANN Transfer Policy](https://www.icann.org/resources/pages/transfer-policy-2016-06-01-en)

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development workflow and commit message conventions.

---

**Investigation Date:** January 30, 2026  
**Status:** ✅ Complete and Ready for Implementation  
**Risk Level:** 🟢 LOW  
**Recommendation:** ✅ Proceed with Implementation
