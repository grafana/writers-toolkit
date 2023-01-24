---
title: UX writing
description: Guidelines on creating text, style, and tone in UI components 
aliases:
  - /docs/writers-toolkit/latest/style-guide/ux-writing/
weight: 500
keywords:
  - Grafana
  - UX writing
---

# UX writing

These guidelines provide guidance on creating text, style, and tone in the different components that make up the UI. They help us build UIs that enhance the user experience, are easy to use, consistent, and inclusive. These guidelines focus on UX writing. For more details on UI elements, refer to the [Grafana Storybook React component library](https://developers.grafana.com/ui/latest/index.html?path=/story/docs-overview-intro--page). 

## Our top tips for writing UI text

Read some of our tips for writing clear and concise UX microcopy. 

### Use a voice that is friendly and not too formal

Be conversational but don't use the voice you use when texting a friend. 
- Talk to our users like you're an engineer casually talking to other engineers. 
- Sound human and show empathy. We're here to help people complete their work and resolve issues. 
- Refrain from using "please" in your UI text.

**Use:**

> Welcome! We’ll guide you through the process of creating your SLOs.

**Don't use:**

> Here is a guide to set up SLOs.

### Be clear and concise 

The more words you use, the more time you waste. 
- Communicate only essential details. 
- Take time to edit what you write to make every word count.

**Use:**

> Save changes 

**Don't use:**

> Would you like to save your changes?

### Address users in second person

Use "you" or "your" as though the UI is speaking to the users. Don't use "please," though.

**Use:**

> Set the target for your SLO.

**Don't use:**

> The target can be set for an SLO.

**Use:**

> Refer to [About Kubernetes Monitoring](/docs/grafana-cloud/kubernetes-monitoring/about-k8s-monitoring/) for details. 

**Don't use:**

> Please refer to [About Kubernetes Monitoring](/docs/grafana-cloud/kubernetes-monitoring/about-k8s-monitoring/) for details. 

### Use active voice

Use active voice to make clear who is performing the action.

**Use:**

> The server receives the query. 

**Don't use:**

> The query is received by the server.

### Use sentence case in UI elements

Capitalize only the first word in the title, the first word in a subheading after a colon, and any proper nouns.

**Use:**

> Create and manage dashboards to visualize your data. 

**Don't use:**

> Create and Manage Dashboards to Visualize Your Data.

### Avoid words created for UI features

Avoid using UI terms when possible.

**Use:**

> In your Grafana Cloud stack, click **Connections**.  

**Don't use:**

> In your Grafana Cloud stack, click the **Connections** button.
 
### Use numerals

The guideline for writing numbers in most mediums is to spell out the numbers one through nine. When writing UI text, it's best to use numerals (1 - 9) because they're easier to parse.

**Use:**

> You have 3 messages.  

**Don't use:**

> You have three messages.

### Skip unnecessary punctuation

Avoid using periods for single sentences in UI elements.

**Use:**

> Search by data source

**Don't use:**

> Search by data source.

Don't use colons after labels.

**Use:**

> Search by type

**Don't use:**

> Search by type:

Use periods for multiple sentences.

**Use:**

> Metrics, Logs, and Traces are billed based on ingestion. For Metrics, we bill based on the number of active series using the 95th percentile during the period. 

**Don't use:**

> Metrics, Logs, and Traces are billed based on ingestion<p></p><p></p>For Metrics, we bill based on the number of active series using the 95th percentile during the period
### Write scannable descriptive text

Using long blocks of descriptive text reduces readability. Write important information first and use short, bulleted lists. Use [headings](#headings) to divide content.

**Use:**

> Before you begin, make sure you have the following items:<ul><li>The kubectl command-line tool. To learn how to install kubectl, refer to [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl).</li><li>(Optional) The helm command-line tool for managing Helm charts. To learn how to install helm, refer to [Installing Helm](https://helm.sh/docs/intro/install/) in the Helm documentation.</li></ul>

**Don't use:**

> Before you begin, make sure you have the kubectl command-line tool available on your local machine. To learn how to install kubectl, refer to [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) in the Kubernetes documentation. You might also need the helm command-line tool for managing Helm charts. To learn how to install helm, refer to [Installing Helm](https://helm.sh/docs/intro/install/) from the Helm documentation.

## Write text for UI elements

The following sections provide writing guidelines for common UI elements:

- Buttons
- Input fields and validation
- Errors
- Alert modals
- Confirm modals
- Tooltips
- Headings
- Links

If you'd like us to add a UI element, refer to the template at the bottom of this topic.

### Buttons

Use buttons when you want users to take actions, such as adding or creating new records or types of information in an existing system.  

- Start button labels with a verb.
- Aim for using one to two words, with a maximum of four words.
- Make button labels descriptive, and tell the user what action will occur if they click it. Rather than using **OK** or **Cancel**, be specific. For example, use **Save** / **Don't save** rather than **Save** / **Cancel**.
- Use sentence case without punctuation.

#### Common use cases for buttons

| Button | When to use it | Examples |
|---|---|---|
| Create | Use when you're creating something new, from scratch or from within an existing app. Creating can involve building something (often large) from scratch. | **Create** <p></p> **Create incident** <p></p> **Create entry** |
| Add | Use when you're adding something into a larger thing. Add extra details to an existing thing. Adding is usually a small or single step, with less work than creating. | **Add details** <p></p> **Add contact** |
| Save | Save something to a server. | **Save** <p></p> **Save and exit** <p></p> **Save changes** |
| Edit | Change or update something that already exists. This doesn’t affect the server until the user saves. | **Edit** |
| Preview | Preview a runtime version of whatever you are working on. This action doesn’t take you away from or override the page you’re already on. | **Preview** |
| Cancel | Cancel or leave a process. Leave without saving any changes. Although **Cancel** is a common button name, it's better to be specific and name the action that's being cancelled.   | **Cancel** |
| Close | Close a window. | **Close** |
| Delete | Permanently delete something from the server. This usually prompts a confirmation modal, asking you to confirm your decision. | **Delete notification policy** <p></p> Deleting this notification policy will permanently remove it. Are you sure you want to delete this policy? <p></p>**Yes, delete** |
| Remove  | Remove an item from a list. | **Remove** |

Refer to the Grafana Storybook React component library for button [usage](https://developers.grafana.com/ui/latest/index.html?path=/docs/buttons-button--basic) and an [example](https://developers.grafana.com/ui/latest/index.html?path=/story/buttons-button--basic).

## Input fields

Use short and scannable text for input field labels. Use sentence case and front-load your field labels with terms that most clearly describe the values they need to enter in the input field. 

If you provide instructions for an input field, be clear about limitations, requirements, and available characters for that field. Use a red asterisk for required fields. Use sentence case without punctuation for the instructions, unless there are multiple sentences.

Optionally, you can provide descriptive placeholder text in an input field. If you do so, make your description clear and concise. The placeholder text should be a hint of the value to be expected. 

**Use:**

> **SLO description** ______________________<p></p>This description appears in a panel next to your SLO status.

**Don't use:**

> **SLO description that appears in a panel next to your SLO status.** <p></p> ______________________

Refer to the Grafana Storybook React component library for button [usage](https://developers.grafana.com/ui/latest/index.html?path=/docs/forms-inlinefield--with-tooltip) and an [example](https://developers.grafana.com/ui/latest/index.html?path=/story/forms-inlinefield--basic).

## Errors

Make errors visible to users, helpful, and easy to understand. Error messages tell the user what happened and what they can do to fix the error.  

- The error headline includes a concise, meaningful summary of the error. Error details provide as much information as possible. 
- Include reasons and instructions for fixing the issue, if possible. Don’t give just the system logs or error titles; try to state how the error was caused and what the user can do to fix it. Assume that some of your users might not understand crash logs and need a simpler description.
- Write error messages for humans without blame. Don't use overdramatic wording or jargon, and avoid apologies. Give a no-nonsense summary of what went wrong and include the degree of severity, in understandable terms.
- Like other UI elements, use sentence case, plain language, and active voice in both the  title and details. Do not use the term "invalid" in an error message. Instead use "not valid" if necessary.

Refer also to [Alert modals](#alert-modals). 

**Use:**

> Failed to evaluate queries and expressions:<p></p>Add a query target to alert rule.

**Don't use:**

> Failed to evaluate queries and expressions: [plugin.downstreamError] failed to query data: no query target found for the alert rule.

**Use:**

> Complete the required fields marked by *.

**Don't use:**

> It looks like you missed a required field.

**Use:**

> The server is unresponsive.

**Don't use:**

> Oops, the server appears to be on a break.


Refer to the Grafana Storybook React component library for input field [usage](https://developers.grafana.com/ui/latest/index.html?path=/docs/forms-input--with-field-validation) and an [example](https://developers.grafana.com/ui/latest/index.html?path=/story/forms-input--simple). 

### Input field validation

Use an input field validation when a text field has formatting requirements. If the validation fails, show the error message directly below the field. 

Refer to the Grafana Storybook React component library for an example of an [input field with validation](https://developers.grafana.com/ui/latest/index.html?path=/story/forms-input--with-field-validation).

## Alert modals

An alert modal displays an important message in a way that attracts the user's attention without interrupting the user's task.

Assume that some of your users might not understand technical terms and need simple, clear alert messages. Like other UI elements, use sentence case, plain language, and active voice in alerts.

### Severity levels

Alert modals have severity levels (error, warning, info, and success) with different colors used for each level:

| Severity | When to use |
|---|---|
| ![Error alert](error.png)| Use an error if an action fails and the user is prevented from completing their task. |
| ![Warning alert](warning.png) | Use a warning to say "don't do this," for example, if the step might be irreversible, leading to permanent data loss. |
| ![Info alert](info.png) | Use as a note to provide useful but not critical information. |
| ![Success alert](success.png) | Use to indicate that an action has completed without errors. |

### Actionable instructions

For error messages, provide actionable instructions to help users complete their task successfully.

**Use:**

> Warning <br> You'll need additional permissions to perform this action. <br>Permissions needed: `plugins:write`

**Don't use:**

> Error <p></p> You'll need additional permissions to perform this action.

## Confirm modals

Use confirm modals to request the user to confirm an action, for example, a deletion. Confirm modals interrupt the user in their flow and force them to deal with the action in the modal. Only use a modal if this interruption is a good thing, for example, when the cost of an error is high. 

- Use affirmative actions with verbs in confirmation messages. Direct and actionable language encourages the user to take the next step. 
- Be sure to also explain the impact and consequences of the options that the user can take. Like other UI elements, use sentence case, plain language, and active voice in the confirmation message title and details.

**Use:**

> Use Google's location service?<p></p>Let Google help apps determine location. This means sending anonymous location data to Google, even when no apps are running.

**Don't use:**

> Are you sure?

Refer to the Grafana Storybook React component library for confirm modal [usage](https://developers.grafana.com/ui/latest/index.html?path=/docs/overlays-confirmmodal--alternative-action) and an [example](https://developers.grafana.com/ui/latest/index.html?path=/story/overlays-confirmmodal--basic).

## Tooltips

Use tooltips to identify UI objects, such as icons. Users hover over a UI object to view a box with a description. 

- Use tooltips for ancillary information since users only refer to the information if they hover over the object. 
- Keep tooltips brief, generally fewer than 120 characters. 
- Consider using tooltips for additional in-app documentation as in this example:
    ![Warning alert](tooltip.png)

Refer to the Grafana Storybook React component library for tooltip modal [usage](https://developers.grafana.com/ui/latest/index.html?path=/docs/overlays-tooltip--basic) and an [example](https://developers.grafana.com/ui/latest/index.html?path=/story/overlays-tooltip--basic).

## Headings

A heading gives structure to your UI elements. Use headings whenever you need to break your content down into hierarchical chunks, often in windows, dialog boxes, and wizards.

- Headings are specific and meaningful and include the most relevant keywords and main points of the chunk, while staying short.
- For headings, use [sentence case](https://developers.google.com/style/capitalization#capitalization-in-titles-and-headings) without punctuation except for question marks when needed. Front-load your headings by putting the word people are looking for at the front of your headline.

**Use:**

> Connect your data to Grafana Cloud

**Don't use:**

> Connect Your Data to Grafana Cloud

## Links

If your product is complex, you might be unable to provide relevant details concisely in the UI text. In this case, you can provide links to documentation for details.  

- Use links sparingly. Try first to write concise and complete UI text. If you include a link, make sure the referenced content will help the user with the task they are completing in the UI.  
- Your link text should be descriptive, telling the user what content they will find upon clicking. Use the exact title of the topic they're linking to so that if the link breaks, they can search for the topic. Like headings, front-load the link text by putting the word people are looking for at the front of your link. 
- Include the link either at the beginning or end of a sentence, not in the middle. Do not include preceding articles as part of the linked text.
- For overview text, you can link to relevant overview documentation and Grafana University courses. 

**Use**

> **Get started** <p></p> Create an alert rule by adding queries and expressions from multiple data sources.<ul><li>Add labels to your alert rules to connect them to notification policies.</li><li>Configure contact points to define where to send your notifications to.</li><li>Configure notification policies to route your alert instances to contact points.</li></ul><p></p>[Read an overview in the documentation >](/docs/grafana/latest/alerting/fundamentals/alert-rules/)<br>[Learn more in the Grafana University course >](https://university.grafana.com/learn/course/external/view/elearning/82/module-intro-to-grafana-alerting)

## Additional elements

The UI elements described here are not exhaustive. If you'd like specific types of UI elements added, let us know. Internal contributors can reach out on Slack and external contributors can reach us using our docs@grafana.com email.  

Use the following template to provide input.

### Template

Use this template to add guidelines for UI elements. 

**[UI element name]**

Write an introduction about the element. Say what its intent or purpose is in an experience.

Provide writing guidelines for the element.

Provide links to usage and examples of the element in the [Grafana Storybook React component library](https://developers.grafana.com/ui/latest/index.html?path=/story/docs-overview-intro--page). 

**Use** 

> Provide an example of how to use the element.

**Don't use**

> Provide an example of how not to use the element.



