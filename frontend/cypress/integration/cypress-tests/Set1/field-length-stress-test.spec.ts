/// <reference types="cypress" />

//Test Objective
//The objective of this test is to check the input for correct size constraints.
//Here we will attempt to inject values of extreamly large sizes and expect the 
//application to reject the request.

describe('Create page', function () {
    it('Access the page', function() {
        cy.visit('http://localhost:4200/create')
        
    })

    it('1st image field', function() {
        const main_image = 'cypress/fixtures/images/elec.jpg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(0)
        .find('input').selectFile(main_image)        
    })
    it('2nd image field', function() {
        const detailed_images = 'cypress/fixtures/images/gadgets.jpeg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(1)
        .find('input').selectFile(detailed_images)        
    })
    it('Form text fields', function() {
        //Long string inputs should be rejected by the application when the filter is set.
        cy.get('input')
        .eq(2)
        .type('Cypress Electronics long string input test Cypress Electronics long string input test Cypress Electronics long string input test Cypress Electronics long string input test Cypress Electronics long string input test Cypress Electronics long string input test Cypress Electronics long string input test Cypress Electronics long string input test Cypress Electronics long string input test ') 
        cy.get('input')
        .eq(3)
        .type('Secondary title Electronics long string input test Secondary title Electronics long string input testSecondary title Electronics long string input testSecondary title Electronics long string input testSecondary title Electronics long string input testSecondary title Electronics long string input testSecondary title Electronics long string input testSecondary title Electronics long string input test') 
        // Numerical input is greater than the Int variable size
        cy.get('input').eq(5).clear()
        .type('152323232320023232323879283687235462354672356745263764738892839209302390283928632763726732837928302930290320302903920362585')
        cy.get('input')
        .eq(6)
        .type('Simple Description long string Simple Description long string Simple Description long string Simple Description long string Simple Description long string v Simple Description long string Simple Description long string Simple Description long string Simple Description long string Simple Description long string Simple Description long string Simple Description long string')
        cy.get('input')
        .eq(7)
        .type('Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string Description for Cypress Iphone 1 long string ')
        cy.get('input')
        .eq(8)
        .type('Mysoreville string test Mysoreville string testMysoreville string testMysoreville string testMysoreville string testMysoreville string testMysoreville string testMysoreville string testMysoreville string test')
        cy.get('input')
        .eq(9)
        .type('Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test Florida string test ')
        cy.get('input')
        .eq(10)
        .type('11241342345145623546723567452637647388928392093023902839286327632675687466235467235674526376473889283920930239028392863276362354672356745263764738892839209302390283928632763623546723567452637647388928392093023902839286327637643422432323232.1')
        cy.get('input')
        .eq(11)
        .type('1623546723567452637647388928392093023902839286327636235467235674526376473889283920930239028392863276362354672356745263764738892839209302390283928632763623546723567452637647388928392093023902839286327635.1')
        cy.get('input')
        .eq(12)
        .clear()
        .type('New')
        cy.get('input')
        .eq(13)
        .clear()
        .type('56235467235674526376473889283920930239028392863276362354672356745263764738892839209302390283928632763623546723567452637647388928392093023902839286327636235467235674526376473889283920930239028392863276362354672356745263764738892839209302390283928632763')
        cy.get('input')
        .eq(15)
        .type('Student string test Student string test Student string test Student string test Student string test Student string test Student string test Student string test Student string test Student string test Student string test Student string test Student string test ')
        
        cy.get('input').eq(16)
        .type('Electronics type string test Electronics type string test Electronics type string test Electronics type string test Electronics type string test Electronics type string test Electronics type string test Electronics type string test Electronics type string test Electronics type string test ')
    })
    it('Submit the form', function() {
        cy.get('button').click()       
    })
})

