Feature: sample karate test script
  for help, see: https://github.com/intuit/karate/wiki/IDE-Support

  Background:
    * url 'http://localhost:3000'

  Scenario: POST1 Test
    Given path 'apipost1'
    When method post
    Then status 200
    And assert 'Post1' == response.message

  Scenario: POST2 Test
    Given path 'apipost2'
    When method post
    Then status 200
    And assert 'Post2' == response.message

  Scenario: POST3 Test
    Given path 'apipost3'
    When method post
    Then status 500
    And assert 'Post3' == response.message

  Scenario: POST4 Test
    Given path 'apipost4'
    When method post
    Then status 403
    And assert 'Post4' == response.message

  Scenario: POST5 Test
    Given path 'apipost5'
    When method post
    Then status 200
    And assert 'Post5' == response.message
  