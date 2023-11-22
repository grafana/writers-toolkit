---
title: Contribute to What's new or release notes
menuTitle: Contribute to What's new or release notes
description: This section describes the different ways of contributing to the What's new document or release notes.
weight: 250
aliases:
  - /docs/writers-toolkit/writing-guide/contribute-release-notes
keywords:
  - what's new
  - release notes
---

# Contribute to What's new or release notes

This topic explains the decisions and actions associated with collecting, writing, and publishing What's new content.

{{% admonition type="note" %}}
This topic is only relevant for internal Grafana Labs contributors.
{{% /admonition %}}

## What's new doc development process

What’s new content is published to the website through the website content management system (CMS), and can be accessed [here](https://admin.grafana.com/content-admin/#/collections/whatsnew).

Because this platform is meant to be used by the entire organization, by default anyone can contribute and publish to What’s new, without the need for approval. **Quality assurance is a conversation within and between contributing teams and internal stakeholders**, but there are some best practice guidelines described in the last two sections of this topic.

Release notes should be entered into the CMS 2-4 weeks before the feature is available, depending on the size of the product or feature. This gives the GTM team time for promotion and enablement. For Grafana versioned releases, have your content entered in the CMS by the cut-off date communicated by the delivery team. For more information, see the [RADS guidelines](https://grafana-intranet--simpplr.vf.force.com/apex/simpplr__app?u=/site/a145f000001dCXBAA2/page/a125f000004oOF7AAM).

When you’re ready to add a What’s new entry, complete the following steps:

1. Fill out the fields:

   - **FEATURE NAME** - Short headline for the improvement. For example, “Grafana OnCall integration for Alerting.”
   - **DATE** - Date and time (UTC) that you want this note published. Typically, this is also the feature release date. If you’ve opened a review PR, you must merge it before the date/time you've added here. If you enter a date that has passed, the note is published immediately on that date.
   - **CONTACT** - First and last name. The contents of this field aren't publicly viewable.
   - **TAGS (OPTIONAL)** - Select category tags that users can use to filter their view. Select as many as apply.
   - **CLOUD AVAILABILITY** - Select the stage of the feature’s Cloud release.
   - **CLOUD OFFERING** - Select which account types have access to the feature.
   - **SELF-MANAGED AVAILABILITY** - Select the stage of the feature's self-managed release.
   - **SELF-MANAGED EDITIONS** - Select the on-premises offerings where this feature is available, or if it's not available in self-managed editions, select "None."
   - **BODY** - Include an overview of the feature and the problem it solves.

     If you want to view some best practices around what to write here, see Writing guidelines for what’s new below.
     Add any images and a link to your public YouTube video here.
     If you need more information on adding an image, refer to [Image, diagram, and screenshot guidelines](https://grafana.com/docs/writers-toolkit/write/image-guidelines/).

   - **FEATURE TOGGLE NAME (OPTIONAL)** - Exact name of the feature toggle for this feature, if one exists. For example, `publicDashboards`.
   - **DOCUMENTATION URL (OPTIONAL)** - URL to the public documentation for this feature.
   - **GRAFANA URL PATH (OPTIONAL)** - URL path to the feature in a Grafana Cloud stack. For example, `/alerting/notifications`.
   - **INTERNAL INFORMATION (OPTIONAL)** - Information for Grafana Labs employees only. For example, ProductDNA, slack channel, FAQ, training docs or videos. Used for training and internal announcements.

1. Click **Save**. The entry is now in **Draft** status.
1. If your entry is ready to publish, proceed to step 4. If your entry requires review, follow these steps:

   1. In the **Status** drop-down, select **In review.**

      A review PR is generated in the `grafana/website` repository. **The Documentation Team does not automatically review these requests; teams that create What’s new entries are responsible for determining their own review process.** However, there are two weekly Office Hours meetings offered by the Documentation Team that you’re welcome to attend for guidance and assistance:

      - [Docs squad office hours (early)](https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=NmoxZ2w0NHBoazgxaW80cTN0MW82aG1xMmxfMjAyMzExMjJUMDkwMDAwWiBpc2FiZWwubWF0d2F3YW5hQGdyYWZhbmEuY29t&tmsrc=isabel.matwawana%40grafana.com&scp=ALL)
      - [Docs squad office hours (late)](https://calendar.google.com/calendar/event?action=TEMPLATE&tmeid=NTI0MWZxYTYwY3FjYW5waWM1OWVjMDBkaWpfMjAyMzExMjJUMjEwMDAwWiBpc2FiZWwubWF0d2F3YW5hQGdyYWZhbmEuY29t&tmsrc=isabel.matwawana%40grafana.com&scp=ALL)

   1. Merge your PR in time for your feature release date.

      Merging your PR ensures your entry is published on the date you entered and it automatically updates the status of your entry in the CMS.

1. To publish your entry from the CMS, follow these steps:

   1. In the **Status** drop-down, click **Ready**.
   1. In the **Publish** drop-down, click **Publish now**. The entry appears in [What's new in Cloud](https://grafana.com/docs/grafana-cloud/whatsnew/) on the date you entered.

For Grafana versioned releases, the content entered in the CMS is published in the versioned What’s new at a later date. Refer to [Creating the self-managed/ppversioned release notes](#create-the-versioned-release-notes)

### Edit What's new entries

Whether your entry is published or not, it's always best to use the CMS to make any changes.

If your entry is published in both _What's new in Cloud_ and _What's new in Grafana_, or it's after the cut-off date for a versioned release, update the CMS and then reach out to the person who's creating the versioned release notes.

### Create the versioned release notes

1. After the previous version of Grafana is released, the DRI cuts a branch and creates a draft PR with an empty What's new doc to be populated with the What's new content for the next release, along with the updated What's new index page. This PR:

   - Should include an update to the link and version number located on the What's new tile of `docs/sources/_index.md`.
   - Should include the new Upgrade Guide page.
   - Should have a `no-backport` label.
   - May include a new Breaking changes guide page.

1. After the cut-off date for the self-managed/versioned release, the DRI goes [where] and filters the What's new content in the CMS by the following properties:

- Self-managed editions
- Date range (from the previous release to the current date)

1. The DRI adds this content to the What's new doc using the tags data to group items appropriately.

1. PM reviews the content to adjust order, **but not copy**. All copy editing must take place when entries are added in the CMS.

1. PM and the DRI work to make final adjustments to the Upgrade guide or Breaking changes guide.

1. A week before the release date, the the DRI changes the PR status from **Draft** to **Ready for Review** to signal to other stakeholders that the PR is now ready for any further review.

1. The DRI finalizes the What's new.

1. The DRI coordinates with the Release Guild and the GTM team for precise timing of when to merge the What's new doc into `main`.

1. On release day, the DRI and merges the What's new branch into `main`. If `main` is no longer the same release as the upcoming release, the DRI should add the appropriate backport label to the PR.

<!-- vale Google.Will = NO -->
<!-- This section speaks of the future -->

The What's new is published in the "next" docs.
When the release branch is promoted to GA, the What's new will also be part of the "latest" doc set.

<!-- vale Google.Will = YES -->

## How to determine if content belongs in What's new

Grafana publishes a What's new [docs page](/docs/grafana/latest/whatsnew/) and [blog post](/blog/2022/11/29/grafana-9.3-release/) along with every minor and major release.

These posts are popular, and a good way for users to learn about the exciting new things Grafana has released.
What's new also drives Grafana's go-to-market enablement: it is used to train the field and make videos on the topics in What's new.

However, unlike a comprehensive changelog, What's new is curated.
If it contained every update and a detailed What's new post for every little bug fix, it would be too noisy for people to read.

So how do you decide whether to write a What's new post for your latest improvement?

### Add a What's new for anything that could excite or concern a customer

What's new content should address changes that have some kind of material impact on the user experience.

- Include changes that _affect customers_, whether they are new features to try out or important, long-requested bug fixes.
  - Most visualization changes and most additions to the UI should be in the What's new document, even when they seem small.
- Almost every change or addition associated with Prometheus and Loki is of interest, too.
- What's new content should also include changes that require customers to do something, like change their API keys to Service Accounts, or stop using a deprecated API or plugin.
- A What's new doc should include announcements—things for customers to notice and try out.
  These could also be notable community contributions to thank a contributor.

When in doubt, ask your nearest PM or EM. Err on the side of _yes, put it in What's new_.

### Examples of what to include in What's new

- A new Transformation: [Partition by values](/docs/grafana/latest/whatsnew/whats-new-in-v9-3/#new-transformation-partition-by-values).
  - This is one of many transformations, but it is brand-new functionality that a user might not notice if they didn't read the What's new document.
    What's new is also a low-effort place to describe some nice use cases and examples for the feature so that users adopt it.
- The new [Candlestick visualization](/docs/grafana/latest/whatsnew/whats-new-in-v8-3/#candlestick-panel-beta).
  - This was a beta feature, but still listed in What's new to get the word out and encourage users to try it.
- All-new Swagger docs for the API.
  - This is significant because it makes Grafana documentation much easier to use, and it's a new place for users to go for help when using the API.
- [Removing beta labels from several panels](https://github.com/grafana/grafana/pull/58557), which makes then generally available.
  - This is a small change code-wise, but a big statement with big customer impact.
    Customers now know that those plugins are fully supported and recommended for use in production.
- [New keyboard shortcut](https://github.com/grafana/grafana/pull/61837)
  - This is a small change, but it brings attention to a feature that has been improved recently and that most people don't know about.
- [Search improvement for Flame graphs](https://github.com/grafana/grafana/pull/61748)
  - Fuzzy search. Has to be in the blog post.
- [Changes to the Prometheus query editor](https://github.com/grafana/grafana/pull/60718)
  - These are query patterns for the data source that most users use.

### Examples of what _not_ to include in What's new

These are important improvements, but are better placed in the changelog than What's new:

- [A docs update](https://github.com/grafana/grafana/pull/60352)
  - This update doesn't require customers to change their behavior—they'll simply see better instructions the next time they use the docs.
- [A bug fix related to migrations](https://github.com/grafana/grafana/pull/59438)
  - This is a bug fix that doesn't require customer action.
- [A usability improvement to an existing transformation](https://github.com/grafana/grafana/pull/59074)
  - Nice fix, but very detailed. Should be in the changelog but not What's new.
- [Change regular expression to accommodate a new branching strategy in Enterprise](https://github.com/grafana/grafana/pull/59429)
  - This change is invisible to customers.

## Writing guidelines for What's new content

Follow these guidelines to ensure that your What's new or release notes content is clear, helpful, and easy to understand.

- **Directly address your users.**

  Address them using the imperative or as “you”.

  **Example:**
  Shorten your communication time when reporting issues and requesting help from Grafana Labs by grabbing a panel's query response data and panel settings.

- **Make use of active voice or present tense.**

  **Example:**
  Enable a configuration option to skip user organization and roles synchronization with your SAML provider.

- **Don't refer to the release version, for example, “In Grafana 9, we ” or “As of now, we”.**

  The What's new or release notes are understood to be providing information for a specific release, so there is no need to repeat this information.

- **Provide high-level descriptions.**

  Tell customers what goal they can accomplish or what problems they can solve with the feature.
  Describe the business value.
  Don't go into details about how the feature works, or configuration steps.

  Link out to topics in the documentation that provide more detailed information or steps.

  **Example:**
  Use custom branding to make Grafana your observability tool by adding your own sign-in page, help links, logo, application name, and more.

  For more information, refer to [insert link to documentation].

- **Don't refer to how the feature used to work.**

  For example, don't say "Previously, alert rules changed state when the rule was facing an error or a timeout.
  Now, the state does not update."

- **For changes or updates to features, provide brief descriptions.**
