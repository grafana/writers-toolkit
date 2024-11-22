---
date: "2023-09-27T17:05:44-04:00"
description: Guidelines for referring to UI elements in Grafana documentation.
keywords:
  - Grafana
  - ui elements
review_date: "2024-05-16"
title: UI elements list
weight: 700
---

# UI elements list

This list provides guidance for how to refer to user interface (UI) elements in Grafana.

{{< admonition type="note" >}}
To view the images in this table, click **Expand Table** at the top-right corner of the table.
{{< /admonition >}}

<!-- prettier-ignore-start -->
<!-- vale Grafana.DialogBox = NO -->
<!-- vale Grafana.WordList = NO -->
<!-- This table includes examples of improper usage of modal, dialog, and hamburger menu. -->

| UI element | Description | Image | Notes |
|------------|-------------|-------|-------|
| breadcrumb | The area at the top of the screen that displays your location in the application and allows you to navigate between the displayed locations. | ![Breadcrumb](/media/docs/writers-toolkit/ui-elements/screenshot-breadcrumb-10.2.png) | Don't refer to this as the _navigation bar_, _top navigation_, or _header_ except in the case of the dashboard header (see below). |
| dashboard header | The area at the top of the screen that displays your location in the application and includes dashboard tools. | ![Dashboard header](/media/docs/writers-toolkit/ui-elements/screenshot-dashboard-header-10.2-2.png) |  |
| dialog box | A box that appears on top of the main screen of the application where you can enter information or select commands. | ![Dialog box](/media/docs/writers-toolkit/ui-elements/screenshot-dialog-box-10.2.png) | Don't use _modal_ or _dialog_ without _box_. |
| drawer | A pane that slides out from the right side of the screen. | ![Sharing drawer](/media/docs/writers-toolkit/ui-elements/screenshot-drawer-11.3.png) | Don't use _panel_, _pop-up_, or _view_. |
| main menu | The navigation that opens on the left side of the screen when you click the logo in top-left corner. | ![Main menu](/media/docs/writers-toolkit/ui-elements/screenshot-main-menu-11.3.png) | Don't refer to this as the _navigation menu_ or _primary menu_. |
| menu icon | The element that opens a menu. | ![Menu icon](/media/docs/writers-toolkit/ui-elements/screenshot-menu-icon-10.2.png) ![Menu icon](/media/docs/writers-toolkit/ui-elements/screenshot-menu-icon-2.png) | Don't refer to this as the _hamburger menu_. |
| panel header | The area at the of a panel that contains the panel name, the edit menu, and other indicators. | ![Panel header](/media/docs/writers-toolkit/ui-elements/screenshot-panel-header-10.2-2.png) |  |
| switch | Used to turn UI features on and off. | ![Switch](/media/docs/writers-toolkit/ui-elements/switch.png) | Don't refer to this as a _toggle_. |
| tab | An element that lets multiple pages exist within one screen. | ![Tab](/media/docs/writers-toolkit/ui-elements/screenshot-tab-10.2.png) | Other UI elements, such as radio button groups ([example 1](/media/docs/writers-toolkit/ui-elements/screenshot-radio-button-group-1-10.2.png), [example 2](/media/docs/writers-toolkit/ui-elements/screenshot-radio-button-group-2-10.2.png)), behave much like tabs, but shouldn't be referred to as tabs; avoid naming these elements, using just the element label instead. |

<!-- vale Grafana.WordList = YES -->
<!-- vale Grafana.DialogBox = YES -->
<!-- prettier-ignore-end -->
