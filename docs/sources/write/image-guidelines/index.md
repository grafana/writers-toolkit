---
aliases:
  - /docs/writers-toolkit/writing-guide/image-guidelines/
  - /docs/writers-toolkit/write/image-guidelines/
date: "2022-07-29T21:26:08-05:00"
description: How to include images and other media in your Grafana Labs documentation.
keywords:
  - image
  - screenshot
  - guideline
menuTitle: Media guidelines
review_date: "2024-06-14"
title: Image, diagram, screenshot, and video guidelines
weight: 750
---

# Image, diagram, screenshot, and video guidelines

At Grafana Labs, media are an essential part of helpful, clear, and concise content.
Follow the standards in this section when creating images for your documents.

## Image and diagram guidelines

Images and diagrams can complement text, enabling readers to quickly grasp a concept and visualize complex processes in a simplified way.
However, because translation services and tools for the visually impaired don't interpret or translate images well, avoid using images to replace text.

### When to use images and diagrams

Consider using images or diagrams when you need to:

- Clarify configurations and settings, such as the architecture for virtual servers
- Define a complex workflow within a Grafana product

Do _not_ include images or diagrams when:

- A workflow is simplistic
- There is no interaction with a Grafana product

### Image and diagram standards

Use the following standards for images and diagrams:

- **Size**: To reduce page-load time, make images and diagrams as small as you can without compromising their usefulness.

  Unless image clarity is critical to understanding, use 65-75 percent quality when saving your image.
  Use the following recommendations when you size images:

  - Horizontal: maximum width 1200px
  - Vertical: maximum height 900px

- **Scope**: Limit the contents of an image to the relevant portion.

  Don't include distracting or unnecessary content and whitespace, such as your browswer window or tabs, or other window decoration.

- **Format**: PNG and SVG are the preferred image formats.

  Use JPG only for photos.

- **Copyright**: You must determine if an image or diagram is protected by copyright.

  If it is, you must obtain permission and acknowledge credit.

- **File name**: Use the naming convention documented in [Media asset file naming conventions](#media-asset-file-naming-conventions).

- **Light mode/Dark mode**: If the page you're adding the image to is white or light colored, use images in dark mode. If the page is black or dark colored, use images in light mode. This ensures sufficient contrast.

#### Personally Identifiable Information (PII)

Make sure to mask, modify, or remove any PII such as passwords, logins, account details, or other information that could compromise security.

#### Alt text and figure captions

Make sure to include [alt text](#alt-text) for every image.
Figure captions are optional.

### Diagram assets

To create diagrams, you need to access the recommended software and download the required icons and stencils.

### Request a diagram from Creative Services

The Design team at Grafana Labs can provide support for diagrams developed by internal Grafana Labs contributors.
Contact them directly using their [design request form](https://grafana-intranet--simpplr.visualforce.com/apex/simpplr__app?u=/site/a145f000001LCBhAAO/page/a125f000001AlBMAA0).

## Screenshot guidelines

When you document a user interface, consider whether you need to include screenshots.
Screenshots can be helpful when text alone can't adequately convey instructions.

Users also like screenshots and find them useful.
However, screenshots are difficult and time-consuming to maintain, and can present problems with translation.

As a result, you should minimize the use of screenshots within your content.

### When to use screenshots

Consider using screenshots when you want to:

- Provide an example of a visualization
- Show panels populated with a query and settings
- Orient users in a complicated or long procedure
- Show complex relationships among drop-down menus, such as those that contain multiple subsets of information and many options available for selection
- Emphasize a feature or a change in the UI

### When not to use screenshots

Do _not_ use screenshots for the following items:

- Simple create operations, such as create a user, a team, an organization, and so on
- Primary or secondary navigation items
- Code samples (instead, show code samples in code blocks)
- Dialog boxes that are easy to understand
- Message text (instead, show message text within the Markdown)
- Progress bars
- Simple pages, such as Wizard pages and Welcome pages
- Tables created in another authoring tool
- A page that's likely to change frequently

### Screenshot alternatives

Only add screenshots to your documentation when necessary.
Instead of screenshots, you can consider being explicit about the user interface elements with which the user interacts.
Add the names of buttons, navigation items, toggles, menus, and so on as they appear on the user interface.

For example, don't include a screenshot to illustrate instructions like, "To add a dashboard, click **Dashboard > New Dashboard**".

### Screenshots in tasks

Place a screenshot below the step that it illustrates.
Don't rely on the screenshot to convey the information or values that the user must enter.
If the user must enter specific information, provide that information in the text of the steps.
However, ensure that the screenshot accurately reflects the directions and values in the step text.

### Screenshot guidelines

Consult the following guidelines when you create screenshots:

- **Size**: The maximum width of a screenshot is 600 pixels.
- **Scope**: Limit the screenshot to just the portion of the user interface that shows the action, and enough surrounding detail to help the user locate the item.
- **Annotations**: To annotate a screenshot, use red, hexadecimal color `FF0000`, arrows and boxes.
- **File name**: Use the naming convention documented in [Media asset file naming conventions](#media-asset-file-naming-conventions).
- **Personal identifiable information (PII)**: Make sure to mask, modify, or remove any PII such as passwords, logins, account details, or other information that could compromise security.
- **Alt text**: Make sure to include [alt text](#alt-text) for every image.

## Video guidelines

Use the following guidelines when you add video content:

- Use videos only from official Grafana sources.
- Use videos that are relevant and specific to the topic.
- Include an introduction to the video that explains what readers can expect to find in the video.

  This is important for accessibility and to save readers from watching videos that aren't important to them.

- Use videos sparingly.

  Video content is much more difficult to update than textual documentation.

The Creative Services team creates videos for blog posts and other content.
They upload these videos to Vimeo or YouTube.
The [Grafana YouTube channel](https://www.youtube.com/@Grafana) has good, short content that highlights features and specific topics.

You can embed these videos in your pages.
To embed YouTube videos, use the [`youtube` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#youtube).
To embed Vimeo videos, use the [`vimeo` shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#vimeo).

## Media asset file naming conventions

The table in this section includes file naming conventions for you to follow when you create image assets.

General rules:

- Use lowercase letters in filenames, always.
- Use dashes between words; don't use spaces or underscores.
- Don't type special characters or punctuation marks in a filename except for periods in product version numbers. For example: `grafana-9.2.release.png`.
- Don't use abbreviations that may cause issues with localization.
  - For example, spell out `database` instead of using `db`.
  - When in doubt, don't use an abbreviation.

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
      <td align="left" valign="top">
        <b>Naming convention:</b> <code>&lt;ASSET TYPE&GT;-&LT;VISUAL DESCRIPTION&GT;-&LT;VERSION, IF APPLICABLE?&gt;.png</code><br>
        <br>
        <b>Examples:</b><br>
        <ul>
          <li><code>screenshot-grafana-loki-uptime-dashboard.png</code></li>
          <li><code>screenshot-grafana-mimir-data-flow-diagram.png</code></li>
          <li><code>screenshot-grafana-9-kubernetes-dashboard.png</code></li>
          <li><code>screenshot-grafana-9.0-kubernetes-dashboard.png</code></li>
          <li><code>screenshot-simple-scalable-test-environment-grafana-loki.png</code></li>
        </ul>
      </td>
    </tr>
    <tr>
      <td align="left" valign="top">Icon</td>
      <td align="left" valign="top">
        Be as descriptive as possible.<br>
        <br>
        For example, use `icon-bar-graph.svg` or `icon-graph-bar.svg` instead of `icon-graph.svg`.<br>
        <b>Naming convention:</b> <code>&lt;ASSET TYPE&GT;-&LT;VISUAL DESCRIPTION?&gt;.svg</code><br>
        <br>
        <b>Examples:</b>
        <ul>
          <li><code>icon-bar-graph.svg</code></li>
          <li><code>icon-prometheus.svg</code></li>
        </ul>
      </td>
    </tr>
    <tr>
      <td align="left" valign="top">Logo</td>
      <td align="left" valign="top">
        When you name Grafana logo files, be sure to include the word "Grafana".<br>
        <br>
        <b>Naming convention:</b> <code>&lt;ASSET TYPE&GT;-&LT;VISUAL DESCRIPTION&GT;-&LT;COLOR + ORIENTATION&GT;.&LT;FILE TYPE?&gt;</code><br>
        <br>
        <b>Examples:</b>
        <ul>
          <li><code>logo-prometheus-full-horizontal.svg</code></li>
          <li><code>logo-grafana-loki-full-horizontal.svg</code></li>
        </ul>
      </td>
    </tr>
    <tr>
      <td align="left" valign="top">Photo</td>
      <td align="left" valign="top">
        <b>Naming convention:</b> <code>&lt;ASSET TYPE&GT;-&LT;VISUAL DESCRIPTION?&gt;.jpg</code><br>
        <br>
        <b>Examples:</b>
        <ul>
          <li><code>photo-raji-on-stage-grafanacon-keynote-2022.jpg</code></li>
          <li><code>photo-grafanacon-team-marketing.jpg</code></li>
          <li><code>photo-headshot-mike-szczys.jpg</code></li>
        </ul>
      </td>
    </tr>
    <tr>
      <td align="left" valign="top">Recording</td>
      <td align="left" valign="top">
        <b>Naming convention:</b> <code>&lt;ASSET TYPE&GT;-&LT;VISUAL DESCRIPTION&GT;.&LT;FILE TYPE?&gt;</code><br>
        <br>
        <b>Example:</b>
        <ul>
          <li><code>gif-grafana-share-playlist.mp4</code></li>
        </ul>
      </td>
    </tr>
  </tbody>
</table>

### Recommended image editors

<!-- vale Grafana.Spelling = NO -->

- Linux: Use [gimp](https://www.gimp.org/)
- macOS: Use [Snagit](https://www.techsmith.com/screen-capture.html)
- Windows: Use [Snagit](https://www.techsmith.com/screen-capture.html)
- Web browser: Use [Photopea](https://www.photopea.com/)

<!-- vale Grafana.Spelling = YES -->

## Where to store media assets

All visual assets are stored in Google Cloud Storage, only accessible to Grafana Labs employees.
You use the [asset upload application](https://admin.grafana.com/upload/) to upload assets.
The following table lists the steps you take to provide the Grafana Labs technical documentation team with the image.

{{< admonition type="note" >}}
Don't store images in the local repository, as it prohibits re-use of the asset in other content.
{{< /admonition >}}

<table>
  <thead>
    <tr>
      <th>If you are a</th>
      <th>Complete these steps</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="left" valign="top">Grafana Labs employee</td>
      <td align="left" valign="top">
        <ol>
          <li>Create a PR against the local repository that includes the Markdown file.</li>
          <li>Navigate to <a href="https://admin.grafana.com/upload/">https://admin.grafana.com/upload/</a>.
          </li>
          <li>Find or create a directory under the media directory. To create a directory, <strong>type the directory name</strong> in the upload input field.</li>
          <li>Use a free image optimizer like <a href="https://tinypng.com/">TinyPNG</a> to decrease the image file size.
          </li>
          <li>Browse and select assets to upload.</li>
          <li>Click <strong>Upload</strong>.</li>
          <li>The asset is available under <code>https://grafana.com/media/</code> in the directory where you uploaded it.</li>
          <li>Click <strong>Copy</strong> to copy the path of the file to your clipboard.</li>
          <li>Add the reference to the Markdown file.<br>
          <br>
          The reference that you add to the Markdown renders the image when you build the Grafana website or a local docs preview.</li>
        </ol>
      </td>
    </tr>
    <tr>
      <td align="left" valign="top">Grafana Labs community contributor</td>
      <td align="left" valign="top">
        <ol>
          <li>Create a PR against the local repository that includes the Markdown file.</li>
          <li>Add the image reference to the Markdown file.<br>
          <br>
          You don't need to test that the image renders in a local build of the documentation. The Grafana technical documentation team reviewing the PR ensures that the image reference works correctly.<br></li>
          <li>Attach the image to the GitHub PR description.</li>
        </ol>
      </td>
    </tr>
  </tbody>
</table>

## Add an image to a Markdown file

To add an image to a Markdown file, insert a reference to the image below the associated step and indent it so that the reference aligns with the step text.

The image reference path conforms to the following convention:

`![<ALT TEXT>](/media/<PATH TO IMAGE>)`

For example:

```
![A table from the console bot showing the Property, Location, and Diff of the breaking changes.](/media/breaking-changes-console-screenshot-1.png)
```

## Test images in a local build

{{< admonition type="note" >}}
This section is relevant to internal Grafana Labs contributors.
{{< /admonition >}}

It's important that you generate a local build of your docs so that you can verify that the path to the image, the image size, and the image placement are correct.

1. Follow the steps in [Where to store media assets](#where-to-store-media-assets).

1. Run `make docs` on your branch and verify that the image appears.
1. If it doesn't render as you desire, upload another version and test again.

## Alt text

Alt text is an HTML attribute that you can use to provide a concise description of an image.
The text is used in situations where the image isn't visible, such as when people use screen readers or have a low-bandwidth internet connection.

In Markdown, the alt text is the text in square brackets when declaring an image:

```markdown
![<ALT TEXT>](/media/<PATH TO IMAGE>)
```

Or, if you're using the `figure` [shortcode](https://grafana.com/docs/writers-toolkit/write/shortcodes/#figure), you can use the `alt` parameter:

```markdown
{{</* figure alt="<ALT TEXT>"  src="/media/<PATH TO IMAGE>" */>}}
```

Every image that you add to the Grafana documentation should have an alt text.

A good piece of guidance for writing good alt text from [HTML: The Living Standard](https://html.spec.whatwg.org/dev/images.html#alt) is:

> One way to think of alternative text is to think about how you would read the page containing the image to someone over the phone, without mentioning that there is an image present. Whatever you say instead of the image is typically a good start for writing the alternative text.

For more information about how to write good alt text, refer to:

- [Google developer documentation style guide](https://developers.google.com/style/images#alt-text)
- [HTML: The Living Standard](https://html.spec.whatwg.org/dev/images.html#alt)
- [gov.uk Design102 blog: What's the alternative? How to write good alt text](https://design102.blog.gov.uk/2022/01/14/whats-the-alternative-how-to-write-good-alt-text/)
- [W3C Web Accessibility Initiative: An alt Decision Tree](https://www.w3.org/WAI/tutorials/images/decision-tree/)

## Screen recordings

The recommended format for screen recordings is `.mp4`.
Don't use `.gif` or `.mov` formats.
In particular, avoid animated GIFs because they distract users.
For more information about distracting content, refer to [Understanding Success Criterion 2.2.2: Pause, Stop, Hide | WAI | W3C](https://www.w3.org/WAI/WCAG21/Understanding/pause-stop-hide.html#intent).

Screen recordings follow the same upload procedure and file naming convention as other media assets.
Use the `video-embed` shortcode to embed the video on the page:

```
{{</* video-embed src="/media/<path-to-recording/recording.mp4>" */>}}
```
