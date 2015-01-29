# ZenPager

A monitor and alert system for small to medium sized companies.


## System Components

- Monitor
- Alert
- Dashboard


### Monitor


Following types of systems monitored
- servers (able to interact with nagios-plugins)
- applications


#### Functionality
- periodically check on servers / applications
- manage what is monitored
- manage what is considered success / failure
- manage 'checks'


### Alert

Makes use of an escalating system of alerting.
Methods
- paging
- email
- external systems (eg: chatrooms)

#### Functionality
- send out alerts
- manage contacts and groups
- manage when alerts are sent
- manage how alerts are sent
- manage who receives alerts


### Dashboard

Current and historical view of Monitor and Alert components

#### Functionality
- receive stats
- view stats
