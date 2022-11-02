---
title: Image, diagram, and screenshot guidelines
menuTitle: Image guidelines
description: How to include images in your documentation.
aliases:
  - /docs/writers-toolkit/latest/writing-guidelines/image-guidelines/
weight: 750
keywords:
  - image
  - screenshot
  - guideline
---

# Image, diagram, and screenshot guidelines

At Grafana Labs, images are an essential part of helpful, clear, and concise content. 
Follow the standards in this section when creating images for your documents.

## Image and diagram guidelines

Images and diagrams can complement text, enabling readers to quickly grasp a concept and visualize complex processes in a simplified way. 
For example, in Grafana blogs, pictures or diagrams can add visual interest and capture the reader's attention. 
However, because translation services and tools for the visually impaired don't interpret or translate images well, avoid using images to replace text.

### When to use images and diagrams

Consider using images or diagrams when you need to:

-  Clarify configurations and settings, such as the architecture for virtual servers
-  Define a complex workflow within a Grafana product

Do *not* include images or diagrams when:

-  A workflow is simplistic
-  There is no interaction with a Grafana product

### Image and diagram standards

Use the following standards for images and diagrams:

-  **Size**: To reduce page-load time, make images and diagrams as small as you can without compromising their usefulness. 
Unless image clarity is critical to understanding, use 65-75 percent quality when saving your image. 
Use the following recommendations when you size images:
   -  Horizontal: maximum width 1200px
   -  Vertical: maximum height 900px

-  **Scope**: Limit the contents of an image to the relevant portion. 
Do not include distracting or unnecessary content and whitespace.
-  **Format**: Because they are lossless, **PNG** and **SVG** are the preferred image formats. 
We do not recommend JPG images because they are lossy and might look blurry. 
-  **Copyright**: Determine if an image or diagram is protected by copyright. 
If it is, you must obtain permission and acknowledge credit.
-  **File name**: Use the naming convention documented in [Image asset file naming conventions](#image-asset-file-naming-conventions).
-  **Personal identifiable information (PII)**: Make sure to mask, modify, or remove any PII such as passwords, logins, account details, or other information that could compromise security.

### Diagram assets

To create diagrams, you need to access the recommended software and download the required icons and stencils.

### Request a diagram from Creative Services

The Design team at Grafana Labs can provide support for diagrams developed by internal Grafana Labs contributors. 
Please contact them directly using their [design request form](https://grafana-intranet--simpplr.visualforce.com/apex/simpplr__app?u=/site/a145f000001LCBhAAO/page/a125f000001AlBMAA0).

## Screenshot guidelines

When you document a user interface, consider whether you need to include screenshots. 
Screenshots can be helpful when text alone cannot adequately convey instructions. 
Users also like screenshots and find them useful. 
However, screenshots are difficult and time-consuming to maintain, and can present problems with translation. 
As a result, you should minimize the use of screenshots within your content.

### When to use screenshots

Consider using screenshots when you want to:

-  Orient users in a complicated or long procedure
-  Show complex relationships among drop-down menus, such as those that contain multiple subsets of information and many options available for selection
-  Emphasize a new feature or a change in the UI
 
### When not to use screenshots

*Do not* use screenshots for the following items:

-  Code samples (show code samples in code blocks)
-  Dialog boxes that are easy to understand
-  Message text (instead show message text within the Markdown)
-  Progress bars
-  Simple pages, such as Wizard pages and Welcome pages
-  Tables created in another authoring tool
-  A page that is likely to change frequently

### Screenshot alternatives

Only add screenshots to your documentation when necessary. 
Instead of providing screenshots, you can consider being explicit about the user interface elements with which the user interacts. 
Add the names of buttons, navigation items, toggles, menus, and so on as they appear on the user interface. 
For example, do not include a screenshot to illustrate simple instructions like, "To add a dashboard, click **Dashboard > New Dashboard**”.

### Screenshots in tasks

Place a screenshot below the step that it illustrates. 
Do not rely on the screenshot to convey the information or values that the user must enter. 
If the user must enter specific information, provide that information in the text of the steps. 
However, ensure that the screenshot accurately reflects the directions and values in the step text.

### Screenshot guidelines

Consult the following guidelines when you create screenshots:

-  **Size**: The maximum width of a screenshot is 600 pixels.
-  **Scope**: Limit the screenshot to just the portion of the user interface that shows the action, and enough surrounding detail to help the user locate the item.
-  **Annotations**: To annotate a screenshot, use red (hexadecimal color **FF0000**) arrows and boxes.
-  **File name**: Use the naming convention documented in [Image asset file naming conventions](#image-asset-file-naming-conventions).
-  **Personal identifiable information (PII)**: Make sure to mask, modify, or remove any PII such as passwords, logins, account details, or other information that could compromise security.

## Image asset file naming conventions

The table in this section includes file naming conventions for you to follow when you create image assets.

General rules:
- Use lowercase letters in filenames, always.
- Use dashes between words; do not use spaces or underscores.
- Do not type special characters or punctuation marks in a file name except for periods in product version numbers (for example, `grafana-9.2.release.png`.
- Don’t use abbreviations that will cause issues with locationization.
  - For example, spell out `database` instead of using `db`, but using an acronym, such as RBAC, is ok.
  - When in doubt, don’t use an abbreviation.
<table>
    <thead>
        <tr>
            <th>Asset type</th>
            <th>Naming convention</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td align="left" valign="top">Screenshot</td>
            <td align="left" valign="top"><b>Naming convention:</b> [asset type]-[visual description]-[version, if applicable].[file type] <br><br><b>Examples:</b><br>
<ul><li>screenshot-grafana-loki-uptime-dashboard.jpg</li>
<li>screenshot-grafana-mimir-data-flow-diagram.jpg</li>
<li>Screenshot-grafana-9-kubernetes-dashboard.jpg (or)</li>
<li>Screenshot-grafana-9.0-kubernetes-dashboard.jpg</li>
<li>simple-scalable-test-environment-grafana-loki.png</li>
</ul>
</td>
        </tr>
        <tr>
            <td align="left" valign="top">Icon</td>
            <td align="left" valign="top">Be as descriptive as possible.<br><br>
            For example, use `icon-bar-graph.svg` or `icon-graph-bar.svg` instead of `icon-graph.svg`.<br>
            <b>Naming convention:</b> [asset type]-[visual description].[file type]<br><br>
            <b>Examples:</b>
<ul><li>icon-bar-graph.svg</li>
<li>icon-prometheus.svg</li>
</ul>
        </tr>
        <tr>
            <td align="left" valign="top">Logo</td>
            <td align="left" valign="top">When you name Grafana logo files, be sure to include the word “grafana”.<br><br>
            <b>Naming convention:</b>[asset type]-[visual description]-[color + orientation].[file type]
<br><br>
            <b>Examples:</b>
<ul><li>logo-prometheus-full-horizontal.svg</li>
<li>Logo-grafana-loki-full-horizontal.svg</li>
</ul>
        </tr>
       <tr>
            <td align="left" valign="top">Photo</td>
            <td align="left" valign="top"><b>Naming convention:</b>[asset type]-[visual description].[file type]
<br><br>
            <b>Examples:</b>
<ul><li>photo-raji-on-stage-grafanacon-keynote-2022.jpg</li>
<li>photo-grafanacon-team-marketing.jpg</li>
<li>Photo-headshot-mike-szczys.jpeg</li>
</ul>
        </tr>
        <tr>
            <td align="left" valign="top">GIF</td>
            <td align="left" valign="top"><b>Naming convention:</b>[asset type]-[visual description].[file type]
<br><br>
            <b>Example:</b>
<ul><li>gif-grafana-share-playlist.gif</li>
</ul>
        </tr>
    </tbody>
</table>

### Recommended image editors

- Linux: Use [gimp](https://www.gimp.org/)
- macOS: Use [Snagit](https://www.techsmith.com/screen-capture.html)
- Windows: TBD.

## Where to store image files
All visual asset images are stored in the [Grafana website repository](https://github.com/grafana/website/tree/master/static/static/img/docs), which is a private repo only accessible to Grafana Labs employees. 
The following table lists the steps you take to provide the Grafana Labs technical documentation team with the image.

> **Note:** Do not store images in the local repository, as it prohibits re-use of the asset for blogs or landing pages.
<table>
    <thead>
        <tr>
            <th>If you are a...</th>
            <th>Complete these steps</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td align="left" valign="top">Grafana Labs employee</td>
            <td align="left" valign="top">
<ol><li>Create a PR against the local repository that includes the Markdown file.
</li>
<li>Add the image reference to the Markdown file.<br><br>
The image reference that you add to the Markdown renders the image when the Grafana website is built. If you want to test that the image appears correctly, refer to <A href="#test-images-in-a-local-build">Test images in a local build</A>.</li>
</ol>
</td>
        </tr>
        <tr>
            <td align="left" valign="top">Grafana Labs community contributor</td>
            <td align="left" valign="top">
<ol><li>Create a PR against the local repository that includes the Markdown file.
</li>
<li>Add the image reference to the Markdown file.<br><br>
We do not expect you to test that the image renders in a local build of the help documentation. The Grafana technical documentation team will ensure that the image reference works correctly.</li><br>
<li>Attach the image to the PR.</li>
</ol>
</td>
        </tr>
    </tbody>
</table>

## Add an image to a Markdown file
To add an image to a Markdown file, insert a reference to the image below the associated step and indent it so that the reference aligns with the step text.

The image reference path conforms to the following convention:
`![<Short description of image>](/static/img/docs/<path-to-image/image-file.png>)`

For example:
1. In the Visualization list, select a visualization type.

   `![Visualization types](/static/img/docs/panel-editor/select-visualization.png)`

## Test images in a local build

> **Note:** This section is relevant to internal Grafana Labs contributors.

It is important that you generate a local build of your docs so that you can verify that the path to the image, the image size, and the image placement are correct. Because images are stored in the Website repo, you must use a `figure` shortcode that renders the image in a local build of the docs.

> **Note:** The following steps only work in the Grafana Cloud docs and  Grafana docs repos.

1. Create a PR against the website repo that contains the image.

   Store image files in the following website repo directory: `static/static/img/docs`. If a relevant sub-directory doesn't exist, you can create it.

1. After the Website team merges the image PR into the website repo, add the following figure shortcode to your docs: 

   `{{</* figure src="[path to the image in the website repo]" */>}}`
   
   **Example:**
   - Path to image stored in the website repo: `static/static/img/docs/grafana-cloud/k8s-node-capacity.png` 
   - Corresponding shortcode: `{{</* figure src="/static/img/docs/grafana-cloud/k8s-node-capacity.png" */>}}`.

1. Run `make docs` on your branch and verify that the image appears.
1. (Optional) Make adjustments to the image in the website repo and test again.
