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


## Installation / Setup

### Database

Setup up the database and run the relevant sql scripts against it.
(TODO have a migration script the runs all the sql scripts)

### Config File

Create a `config.toml` file at the root of the project. For example;

   [postgresql]
   name = "zenpager"
   user = "postgres"
   password = ""
   host = "localhost"
   sslmode = "disable"

### Create user

Run the following command to create a user in the system;

   make createuser


## Run

To run the application, use the following command;

   make run
