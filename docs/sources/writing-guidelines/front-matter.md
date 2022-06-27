---
title: "Front matter"
menuTitle: "Front matter"
description: “Learn about how we build front matter to properly enable the publication and search of our technical documentation”
aliases: ["/docs/writers-toolkit/latest/writing-guidelines/front-matter"]
weight: 100
Keywords:
    - front matter
    - alias
---

# Front matter

Grafana technical documentation includes front matter to help organize the content, develop the TOC (as published in the LH sidebar of the website), and help users identify useful pages when searching or viewing the content in search engines or in social media, such as Twitter.

We use YAML for all front matter.

Here’s a correctly built example (note that few to none of our files yet use this):

    ---
    title: "About Grafana Mimir architecture"
    menuTitle: "Architecture"
    description: “Learn more about Grafana Mimir’s microservices-based architecture”
    aliases: ["/docs/mimir/latest/old-architecture/"]
    weight: 100
    Keywords:
        - Mimir
        - microservices
        - architecture
    ---

| Element  | Description | Guideline  |
|---|---|---|
| [title] | Required. </br></br>The [title] displays as the H1 on the page. </br></br>The [title] becomes the document `<title>` element. Often browsers display this in the tab for the page.  | Does not need to precisely match the menuTitle. The title should be optimized for search engines. |
| [seo_title] | TBD | TBD |
| [menuTitle] | The [menuTitle] is useful for having a distinct sidebar entry perhaps in the case that the title is too long to display nicely in the sidebar as the title on the website in the LH sidebar.</br></br>Note: Not all repos support [menuTitle], per the website team. | Does not need to precisely match the title. The menuTitle does not need to be optimized for search engines. |
| [description] | The [description] text displays as a clue to users about what the page should include on social (Twitter and the like), though not as much by a search engine. | The number of characters vary by media, but use them wisely. </br></br>Provide enough information to guide users to the content by describing what content is provided using the link. Often, this doesn’t need to be original prose - you can often scan the first few paragraphs to pluck the appropriate terms/phrases into the description. </br></br>It won't cause harm if it's too long, it will simply truncate in the displayed media. |
| [aliases] | Provides an HTML redirect from the pages in the list to the current page. Described in detail in https://github.com/grafana/technical-documentation/blob/main/docs/hugo/index.md  |   |
| [weight] | The [weight] determine the placement of the topic within the left hand sidebar of our website, with smaller numbers placing the topic higher in the guide. </br></br>Pages with the same weight have lexicographic ordering. | We recommend that you use increments of `100` for index files and for all other content files, because doing so eliminates much of the need to re-order existing topics when new topics are added. </br></br>Weights are per web directory. |
| [keywords] | Keywords are used by the website to link to related pages in the “related content” sections. https://github.com/grafana/website/blob/master/config/_default/config.yaml#L85 </br></br>They do not appear in the resulting HTML source for the page and have no effect on SEO. </br></br>Ideally, use single terms as opposed to phrases. |   |
