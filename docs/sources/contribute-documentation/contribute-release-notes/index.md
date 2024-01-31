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

What’s new content is published to the website through the website content management system (CMS), and can be accessed [here](https://admin.grafana.com/content-admin/#/collections/whats-new).

Because this platform is meant to be used by the entire organization, by default anyone can contribute and publish to What’s new, without the need for approval. **Quality assurance is a conversation within and between contributing teams and internal stakeholders**, but there are some best practice guidelines described in the last two sections of this topic.

Release notes should be entered into the CMS 2-4 weeks before the feature is available, depending on the size of the product or feature. This gives the GTM team time for promotion and enablement. For Grafana versioned releases, have your content entered in the CMS by the cut-off date communicated by the delivery team. For more information, see the [RADS guidelines](https://grafana-intranet--simpplr.vf.force.com/apex/simpplr__app?u=/site/a145f000001dCXBAA2/page/a125f000004oOF7AAM).

Note that in the context of the CMS, the word "published" has a slightly different meaning than in general use:

- **Published** - Your entry is complete and in **Published** status. It's either visible on the external _What's new in Cloud_ page or is going to automatically become visible on the release date.
- **Live** - Your entry is visible on the _What's new in Cloud_ page.

### Create a What's new entry

When you’re ready to add a What’s new entry, complete the following steps:

1. Fill out the fields:

   - **FEATURE NAME** - Short headline for the improvement. For example, “Grafana OnCall integration for Alerting.”
   - **FEATURE RELEASE DATE** - Date and time (UTC) that you want this note to be live. This should also be the feature release date. If the feature is behind a feature toggle and gets rolled out only to a fraction of users the date is when the feature was first available to users (opting in).

     If you’ve opened a review PR, you must merge it before the date/time you've added here. If you enter a date that has passed, the note is published immediately with that date.

   - **CONTACT** - First and last name. The contents of this field aren't publicly viewable.
   - **TAGS (OPTIONAL)** - Select category tags that users can use to filter their view. Select as many as apply.
   - **CLOUD AVAILABILITY** - Select the stage of the feature’s Cloud release. If the feature isn't available in Cloud, select "None."
   - **CLOUD EDITIONS** - Select which account types have access to the feature. If the feature isn't available in Cloud, select "None."
   - **SELF-MANAGED AVAILABILITY** - Select the stage of the feature's self-managed release. If the feature isn't available in self-managed Grafana, select "None."
   - **SELF-MANAGED EDITIONS** - Select the on-premises offerings where this feature is available. If the feature isn't available in self-managed Grafana, select "None."
   - **BODY** - Include an overview of the feature and the problem it solves.

     If you want to view some best practices around what to write here, see Writing guidelines for what’s new below.
     Add any images and a link to your public YouTube video here.
     If you need more information on adding an image, refer to [Image, diagram, and screenshot guidelines](https://grafana.com/docs/writers-toolkit/write/image-guidelines/).

   - **DOCUMENTATION URL (OPTIONAL)** - URL to the public documentation for this feature.
   - **ENABLEMENT VIDEO (OPTIONAL)** - Link to the video used for enablement.
   - **INTERNAL INFORMATION (OPTIONAL)** - Information for Grafana Labs employees only. For example, ProductDNA, slack channel, FAQ, training docs or videos. Used for training and internal announcements. This will only be visible on admin.grafana.com/whats-new/, _not_ the public What's New page.

1. Click **Save**. The entry is now in **Draft** status and a PR is generated in the `grafana/website` repository.
1. If your entry is ready to publish, proceed to step 4. If your entry requires review, follow these steps:

   1. In the **Status** drop-down, select **In review.**

      **The Documentation Team does not automatically review these pull requests; teams that create What’s new entries are responsible for determining their own review process.** However, there are two weekly Office Hours meetings offered by the Documentation Team that you’re welcome to attend for guidance and assistance:

      - [Docs squad office hours (early)](https://docs.google.com/document/d/1QaH9PwIZ_6-6Udhmy7Zhwme72ZZLqSTiX8HYFFi6AE4/edit#heading=h.9o53ccbx7xrw)
      - [Docs squad office hours (late)](https://docs.google.com/document/d/12XK3XYEZWU3uIPluq3sn5HSGqWHBHJkktqlyQZHj_Uw/edit#heading=h.9o53ccbx7xrw)

   1. Work with your team to review and finalize the generated pull request.

   1. Merge your PR in time for your feature release date.

      Merging your PR ensures your entry is live on the date you entered and it automatically updates the status of your entry in the CMS.

1. To publish your entry from the CMS, follow these steps:

   1. In the **Status** drop-down, click **Ready**.
   1. In the **Publish** drop-down, click **Publish now**. The entry appears in [What's new in Cloud](https://grafana.com/docs/grafana-cloud/whats-new/) on the date you entered.

For Grafana versioned releases, the content entered in the CMS is published in the versioned What’s new at a later date. Refer to [Creating the versioned release notes](#create-the-versioned-release-notes)

If you add an entry to the CMS after the relevant versioned What's new has already been published, you'll need to open a PR to also add it to the versioned What's new yourself.

### Edit What's new entries

Regardless of the status of your entry, it's always best to use the CMS to make any changes. To make edits, follow these steps:

1. Navigate to the [CMS](https://admin.grafana.com/content-admin/#/collections/whats-new).

   - If your entry is already in **Published** status, you can find it [in the Contents section of the CMS under What's New in Cloud](https://admin.grafana.com/content-admin/#/collections/whats-new).
   - If your entry is still in **Draft** or **Review** status, you can find it [in the Workflow section of the CMS](https://admin.grafana.com/content-admin/#/workflow) under the appropriate heading.

1. Update the fields that you need to change.
1. Click **Save**. The entry is now in **Draft** status.
1. Do one of the following:

   - If your entry is ready to publish, select **Ready** in the **Status** drop-down, and then **Publish now** in the **Publish** drop-down.
   - If your entry needs to be reviewed, select **In review** in the **Status** drop-down to open a review PR. For more information on managing review PRs, see step 3 in the [Create a What's new entry](#create-a-whats-new-entry).

If your entry is already live in both _What's new in Cloud_ and it's between the cut-off date for a versioned release and the release date, update the CMS and then reach out to the person responsible for creating the versioned release notes.

If your entry requires an update after it's live in both the Cloud and self-managed What's new, you'll need to update both entries.

### Create the versioned release notes

The following instructions are for the person directly responsible for creating the versioned release notes. This is typically someone on the Technical Writing team.

1. After the cut-off date for What's new entries has passed, cut a branch and create a draft PR with an empty `whats-new-in-vxx-x.md` file to be populated with the What's new content for the next release. This PR should include:

   - Updates to the `whatsnew/_index.md`
   - Update to the link and version number located on the What's new tile of `docs/sources/_index.md`
   - The new upgrade guide
   - The new breaking changes page, if needed

1. Label the PR `no-backport` for now; this may change.

1. Have someone, typically the Tech Writing team build engineer, generate a Markdown file from the _What's new in Cloud_ with the following conditions:

   - Filtered by the relevant Grafana version
   - Includes front matter for each entry
   - Grouped by tags; entries with multiple tags should only be included once, grouped by their first tag alphabetically

1. Add the content of this Markdown file to the `whats-new-in-vxx-x.md` file using the tags data to group items.

   If internal enablement videos are listed for entries, but the associated YouTube videos aren't in the body text of those entries yet, you'll need to add them later.
   To do this, generate another Markdown file from the _What's new in Cloud_ closer to the release date and make updates in `whats-new-in-vxx-x.md` from the newly generated file.

1. A week before the release date, change the PR status from **Draft** to **Ready for Review** to signal to other stakeholders that the PR is now ready for any further review.

   Reviews _must not include any copy edits_ unless there are inaccuracies or typos, because all copy edits should happen when entries are added in the _What's new in Cloud_.
   If there are any inaccuracies, those need to be corrected in both the Cloud and versioned What's new.

1. Have the PM review the content to adjust the order, if needed.

1. Work with the PM to make final adjustments to the upgrade guide or breaking changes page.

1. Add a backport label to the PR, if needed.

1. Two days before the release, get a final generated markdown file from the _What's new in Cloud_ and make any needed additions to the `whats-new-in-vxx-x.md` file.

1. On the day before release day, merge the What's new branch into `main`.

<!-- vale Google.Will = NO -->
<!-- This section speaks of the future -->

The What's new is live in the "next" docs.
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
