---
aliases:
  - /docs/writers-toolkit/contribute-release-notes/
  - /docs/writers-toolkit/contribute/release-notes/
  - /docs/writers-toolkit/writing-guide/contribute-documentation/contribute-release-notes/
date: 2024-02-14
description: This section describes the different ways of contributing to the What's new document or release notes.
keywords:
  - what's new
  - release notes
menuTitle: Contribute to What's new or release notes
title: Contribute to What's new or release notes
weight: 250
---

<!-- vale Grafana.Timeless = NO -->

# Contribute to What's new or release notes

This topic explains the decisions and actions associated with collecting, writing, and publishing What's new content.

{{< admonition type="note" >}}
This topic is only relevant for internal Grafana Labs contributors.
{{< /admonition >}}

## What's new documentation development process

What’s new content is published to the website through the website Content Management System (CMS).
To add a new note, browse to the [What's new CMS collection](https://admin.grafana.com/content-admin/#/collections/whats-new).

Because this platform is meant to be used by the entire organization, by default anyone can contribute and publish to What’s new, without the need for approval.
_Quality assurance is a conversation within and between contributing teams and internal stakeholders_, but there are some best practice guidelines described in the last two sections of this topic.

Enter release notes into the CMS two to four weeks before the feature is available, depending on the size of the product or feature.
This gives the Go To Market (GTM) team time for promotion and enablement.
For Grafana versioned releases, have your content entered in the CMS by the cut-off date communicated by the delivery team.
For more information, refer to the [Record Announce Document Ship (RADS) guidelines](https://grafana-intranet--simpplr.vf.force.com/apex/simpplr__app?u=/site/a145f000001dCXBAA2/page/a125f000004oOF7AAM).

It's important to understand that, in the context of the CMS, the word "published" has a slightly different meaning than in general use:

- **Published**: Your entry is complete and in **Published** status.
  It's either visible on the external _What's new in Cloud_ page or is going to automatically become visible on the release date.
- **Live**: Your entry is visible on the _What's new in Cloud_ page.

### Create a What's new entry

When you’re ready to add a What’s new entry, complete the following steps:

1. Fill out the fields:
   <table>
     <thead>
       <tr>
         <th scope="col">Field</th>
         <th scope="col">Description</th>
         <th scope="col">Guidance</th>
       </tr>
     </thead>
     <tbody>
       <tr>
         <th scope="row">FEATURE NAME</th>
         <td>Short headline for the feature.</td>
         <td>For example, <em>Grafana OnCall integration for Alerting</em>.</td>
       </tr>
       <tr>
         <th scope="row">FEATURE RELEASE DATE</th>
         <td>Date and time in UTC that you want this note to be live.</td>
         <td>
           <p>This should also be the feature release date. If the feature is behind a feature toggle and gets rolled out only to a fraction of users the date is when the feature was first available to users opting in.</p>
           <p>If you’ve opened a review PR, you must merge it before the date you've added here. If you enter a date that has passed, the website publishes the note on the next build.</p>
         </td>
       </tr>
       <tr>
         <th scope="row">CONTACT</th>
         <td>First and last name.</td>
         <td>The contents of this field aren't publicly viewable.</td>
       </tr>
       <tr>
         <th scope="row">TAGS (OPTIONAL)</th>
         <td>Select category tags that users can use to filter their view.</td>
         <td>Select as many as apply.</td>
       </tr>
       <tr>
         <th scope="row">CLOUD AVAILABILITY</th>
         <td>Select the stage of the feature’s Cloud release.</td>
         <td>If the feature isn't available in Cloud, select <strong>None</strong>.</td>
       </tr>
       <tr>
         <th scope="row">CLOUD EDITIONS</th>
         <td>Select which account types have access to the feature.</td>
         <td>If the feature isn't available in Cloud, select <strong>None</strong>.</td>
       </tr>
       <tr>
         <th scope="row">SELF-MANAGED AVAILABILITY</th>
         <td>Select the stage of the feature's self-managed release.</td>
         <td>If the feature isn't available in the self-managed product, select <strong>None</strong>.</td>
       </tr>
       <tr>
         <th scope="row">SELF-MANAGED EDITIONS</th>
         <td>Select the on-premises offerings where this feature is available.</td>
         <td>If the feature isn't available in the self-managed product, select <strong>None</strong>.</td>
       </tr>
       <tr>
         <th scope="row">SELF-MANAGED VERSION</th>
         <!-- vale Grafana.GoogleWill = NO -->
         <td>Select the version of self-managed product that will include the feature.</td>
         <!-- vale Grafana.GoogleWill = YES -->
         <td>
           <p>If the feature isn't available in the self-managed product, select <strong>None</strong>.</p>
           <p>If the version isn't available, select <strong>No suitable option</strong> and reach out in the #docs Slack channel so that a maintainer can add a new option.
           </p>
         </td>
       </tr>
       <tr>
         <th scope="row">BODY</th>
         <td>Include an overview of the feature and the problem it solves.</td>
         <td>
           <p>If you want to view some best practices around what to write here, refer to <a href="#writing-guidelines-for-whats-new-content">Writing guidelines for What’s new</a>.</p>
           <p>Add any images and a link to your public YouTube video here.</p>
           <p>If you need more information on adding an image, refer to <a href="https://grafana.com/docs/writers-toolkit/write/image-guidelines/">Image, diagram, and screenshot guidelines</a>.</p>
           <p>If you need to mention a feature flag, use this format: To try out Trace to profiles, enable the <code>traceToProfile</code> <a href="https://grafana.com/docs/grafana/latest/setup-grafana/configure-grafana/feature-toggles/">feature toggle</a>.</p>
         </td>
       </tr>
       <tr>
         <th scope="row">DOCUMENTATION URL (OPTIONAL)</th>
         <td>URL to the public documentation for this feature.</td>
         <td></td>
       </tr>
       <tr>
         <th scope="row">ENABLEMENT VIDEO (OPTIONAL)</th>
         <td>Link to the video used for enablement.</td>
         <td>
           <p>Enablement videos are perhaps the fastest, most engaging tool for employees and users to learn about your feature. Use them for maximum engagement.</p>
           <p>Follow these instructions to create and upload a video: <a href="https://docs.google.com/document/d/1nCiG62FxJ9J_qLTnlzNopSsT-VEAlxCUCsrAEpWL68U/edit#heading=h.fierz9i4q8ft">Enablement video instructions</a>.</p>
           <p>When you upload an enablement video, the Content team receives a notification, edits it for the public, and uploads it to YouTube to coincide with your feature's release. They need a few weeks' lead time for this.</p>
         </td>
       </tr>
       <tr>
         <th scope="row">INTERNAL INFORMATION (OPTIONAL)</th>
         <td>Information for Grafana Labs employees only.</td>
         <td>
           <p>For example, ProductDNA, Slack channel, FAQ, training documentation, or videos.</p>
           <p>Used for training and internal announcements.</p>
           <p>This is only visible on the <a href="https://admin.grafana.com/whats-new/">internal What's new page</a>, <em>not</em> the <a href="https://grafana.com/docs/grafana-cloud/whats-new/">public What's New page</a>.</p>
         </td>
       </tr>
     </tbody>
   </table>

1. Click **Save**.
   The entry is now in **Draft** status and the CMS opens a pull request in the `grafana/website` repository.

   If your entry is future-dated, it won't show up in the website-generated preview or in your local build, but you can see it in the preview of the internal feed. To do so, remove `/docs/grafana-cloud` from the preview URL and add the heading for your entry at the end of the URL. It should look something like this: `https://deploy-preview-18347-zb444pucvq-uw.a.run.app/whats-new/#create-subtables-in-table-visualizations`.

1. If your entry is ready to publish, proceed to step 4.
   If your entry requires review, follow these steps:

   1. In the **Status** drop-down, select **In review.**

      {{< admonition type="note" >}}
      _The Documentation Team doesn't automatically review these pull requests; teams that create What’s new entries are responsible for determining their own review process._

      However, there are two weekly Office Hours meetings offered by the Documentation Team that you’re welcome to attend for guidance and assistance:

      - [Docs squad office hours (early)](https://docs.google.com/document/d/1QaH9PwIZ_6-6Udhmy7Zhwme72ZZLqSTiX8HYFFi6AE4/edit#heading=h.9o53ccbx7xrw)
      - [Docs squad office hours (late)](https://docs.google.com/document/d/12XK3XYEZWU3uIPluq3sn5HSGqWHBHJkktqlyQZHj_Uw/edit#heading=h.9o53ccbx7xrw)

      {{< /admonition >}}

   1. Work with your team to review and finalize the generated pull request.

   1. Merge your PR in time for your feature release date.

      Merging your PR ensures your entry is live on the date you entered and it automatically updates the status of your entry in the CMS.

1. To publish your entry from the CMS, follow these steps:

   1. In the **Status** drop-down, click **Ready**.
   1. In the **Publish** drop-down, click **Publish now**.
      The entry appears in [What's new in Cloud](https://grafana.com/docs/grafana-cloud/whats-new/) on the date you entered.

For Grafana versioned releases, the content you enter in the CMS is published in the versioned What’s new at a later date.
To understand the process of creating release notes for Grafana versioned releases, refer to [Create the versioned release notes](#create-the-versioned-release-notes).

If you add an entry to the CMS after the relevant versioned What's new has already been published, you'll need to open a PR to also add it to the versioned What's new yourself.

### Edit What's new entries

Regardless of the status of your entry, it's always best to use the CMS to make any changes.
To make edits, follow these steps:

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

The following instructions are for the person directly responsible for creating the versioned release notes.
This is typically someone on the Technical Writing team.

1. After the cut-off date for What's new entries has passed, cut a branch and create a draft pull request with an empty `whats-new-in-vxx-x.md` file to be populated with the What's new content for the next release.
   This PR should include:

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

1. Two days before the release, get a final generated Markdown file from the _What's new in Cloud_ and make any needed additions to the `whats-new-in-vxx-x.md` file.

1. On the day before release day, merge the What's new branch into `main`.

<!-- vale Grafana.GoogleWill = NO -->
<!-- This section speaks of the future -->

The What's new is live in the "next" docs.
When the release branch is promoted to GA, the What's new will also be part of the "latest" doc set.

<!-- vale Grafana.GoogleWill = YES -->

## How to determine if content belongs in What's new

Grafana publishes a What's new [documentation page](https://grafana.com/docs/grafana/latest/whatsnew/) and [blog post](/blog/2022/11/29/grafana-9.3-release/) along with every minor and major release.

These posts are popular, and a good way for users to learn about the exciting new things Grafana has released.
What's new also drives Grafana's go-to-market enablement: it's used to train the field and make videos on the topics in What's new.

However, unlike a comprehensive changelog, What's new is curated.
If it contained every update and a detailed What's new post for every little bug fix, it would be too noisy for people to read.

So how do you decide whether to write a What's new post for your latest improvement?

### Add a What's new for anything that could excite or concern a customer

What's new content should address changes that have some kind of material impact on the user experience.

- Include changes that _affect customers_, whether they're new features to try out or important, long-requested bug fixes.
  - Most visualization changes and most additions to the UI should be in the What's new document, even when they seem small.
- Almost every change or addition associated with Prometheus and Loki is of interest, too.
- What's new content should also include changes that require customers to do something, like change their API keys to Service Accounts, or stop using a deprecated API or plugin.
- A What's new page should include announcements —- things for customers to notice and try out.
  These could also be notable community contributions to thank a contributor.

When in doubt, ask your nearest Product Manager (PM) or Engineering Manager (EM).
Err on the side of _yes, put it in What's new_.

### Examples of what to include in What's new

- A new Transformation: [Partition by values](/docs/grafana/latest/whatsnew/whats-new-in-v9-3/#new-transformation-partition-by-values).
  - This is one of many transformations, but it's also brand-new functionality that a user might not notice if they didn't read the What's new document.
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

- [A documentation update](https://github.com/grafana/grafana/pull/60352)
  - This update doesn't require customers to change their behavior—they'll simply see better instructions the next time they use the docs.
- [A bug fix related to migrations](https://github.com/grafana/grafana/pull/59438)
  - This is a bug fix that doesn't require customer action.
- [A usability improvement to an existing transformation](https://github.com/grafana/grafana/pull/59074)
  - Nice fix, but very detailed.
    Should be in the changelog but not What's new.
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
  Enable a configuration option to skip user organization and roles synchronization with your Security Assertion Markup Language (SAML) provider.

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
  Now, the state doesn't update."

- **For changes or updates to features, provide brief descriptions.**
