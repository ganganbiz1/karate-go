Feature: sample karate test script
  for help, see: https://github.com/intuit/karate/wiki/IDE-Support

  Background:
    * url 'http://localhost:3000'

  Scenario: GET1 Test
    Given path 'api1'
    When method get
    Then status 200
    And assert 'Get1' == response.message

  Scenario: GET2 Test
    Given path 'api2'
    When method get
    Then status 400
    And assert 'Get2' == response.message

  Scenario: GET3 Test
    Given path 'api3'
    When method get
    Then status 200
    And assert 'Get3' == response.message

  Scenario: GET4 Test
    Given path 'api4'
    When method get
    Then status 200
    And assert 'Get4' == response.message

  Scenario: GET5 Test
    Given path 'api5'
    When method get
    Then status 200
    And assert 'Get5' == response.message
  