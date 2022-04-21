/// <reference types="cypress" />

//Test Objective
//The create page will be accessed and images from the cypress 
//fixtures are loaded and a new product is created.

describe('Create a product test', function () {
    it('Access the page', function() {
        cy.visit('http://localhost:4200/create')
        
    })

    it('Interacting with the 1st image field', function() {
        //adding the relative path - get the images for the fixtures
        const main_image = 'cypress/fixtures/images/elec.jpg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(0)
        .find('input').selectFile(main_image)        
    })
    it('Interacting with the 2nd image field', function() {
        const detailed_images = 'cypress/fixtures/images/gadgets.jpeg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(1)
        .find('input').selectFile(detailed_images)        
    })
    it('Form text fields', function() {
        //Title
        cy.get('input')
        .eq(2)
        .type('Cypress Electronics') 
        //Seconday Title
        cy.get('input')
        .eq(3)
        .type('Secondary title Electronics') 
        //Image url 
        //cy.get('input').eq(4)
        //Price
        cy.get('input')
        .eq(5)
        .clear()
        .type('1500')
        //Simple Description
        cy.get('input')
        .eq(6)
        .type('Simple Description for Cypress Iphone 1')
        //Description
        cy.get('input')
        .eq(7)
        .type('Description for Cypress Iphone 1')
        //City
        cy.get('input')
        .eq(8)
        .type('Mysoreville')
        //State
        cy.get('input')
        .eq(9)
        .type('Florida')
        //Latitude
        cy.get('input')
        .eq(10)
        .type('12.1')
        //Longitude
        cy.get('input')
        .eq(11)
        .type('15.1')
        //Condition
        cy.get('input')
        .eq(12)
        .clear()
        .type('New')
        //Age
        cy.get('input')
        .eq(13)
        .clear()
        .type('5')
        //Image url 
        //cy.get('input').eq(14)
        //Target
        cy.get('input')
        .eq(15)
        .type('Student')
        //Category
        cy.get('input')
        .eq(16)
        .type('Electronics')
    })
    it('Submit the form', function() {
        cy.get('button').click()       
    })

})

