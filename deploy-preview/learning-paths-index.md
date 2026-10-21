---
menuTitle: Paths
title: Paths
weight: 100
type: learning-hub
layout: type
learn_type: path
learn_microsite_chrome: true
featured_items:
  - grafana-cloud-tour
  - drilldown-metrics
  - send-logs-alloy-loki
cascade:
  - _target:
      path: /docs/learning-paths/**
    learn_microsite_chrome: true
  - layout: single-journey
    outputs: [HTML, json, pathfinder, unstyled]
    meta_image: /media/docs/learning-journey/linux/learningjourney-metaimage.png
    GRAFANA_VERSION: next
    search_section: Grafana Labs Learning
    search_type: doc
    meta_icon: learning-paths
hero:
  content: |
    Welcome to Grafana Labs paths. At Grafana Labs, our top priority is your success, and we're dedicated to supporting you every step of the way.

    A path offers a clear and organized route to mastering our platform. Each step includes expert guidance based on what we know works well. You also have the flexibility to browse related side journeys along the way.

    Whether you are an administrator managing thousands of users or a hobbyist, our paths provide an effective way to quickly gain value.

    Select one of our journeys to start your path toward Grafana mastery!
  background_image: '/media/docs/learning-journey/map-background-2.jpg'
  # logo: '/media/docs/learning-journey/map.svg'
  video: 5B-vPhqDib8
map:
  title: New to Grafana Cloud?
  subhead: Consider taking the following path to get started.
journey_categories:
  - id: data-availability
    label: Data availability
  - id: query-and-visualize
    label: Query and visualize
  - id: take-action
    label: Take action
cta:
  defaults:
    success:
      title: Were you successful?
      wasSuccessful:
        title: 'Yes'
        text: Super! You can move to the next step.
      notSuccessful:
        title: 'No'
        text: Need more support? Check out these
        link:
          text: helpful resources.
          href: /help/
        troubleshooting:
          title: Troubleshooting options
        form:
          text: "Please tell us what didn't work:"
          submit: Submit
          loading: Submitting...
          error: There was an error with your submission.
          success: Thank you!
conclusion:
  rating: How satisfied were you with your path?
  message: Thanks for your rating! We'd appreciate any feedback on how we can improve.
  loading: Submitting...
  error: There was an error with your submission.
  success: Thank you!
  submit: Submit
  ratings:
    - Very dissatisfied
    - Somewhat dissatisfied
    - Neither satisfied nor dissatisfied
    - Somewhat satisfied
    - Very satisfied
version_archive_exclude: true
---

# Paths

Paths are short, task-based guides, each focused on a single outcome — deploying a Helm chart, connecting a data source, tuning an alert. Every path is a hands-on sequence you can finish in one sitting, and many let you practice directly in Grafana Play or your own Grafana stack.
