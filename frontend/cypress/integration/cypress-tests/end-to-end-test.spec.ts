// Objective: Complete end to end test Cypress video
/// <reference types ="cypress" />

describe('Register an user', function () {
    it('Sign up', function() {
        cy.visit('http://localhost:4200/signup')
        
    })

    it('Sign up form', function() {
        //Email
        cy.get('input')
        .eq(0).type('newuser@gatormart.com') 
        //Password
        cy.get('input')
        .eq(1).type('Password2$') 
        //Re-Type Password
        cy.get('input')
        .eq(2).type('Password2$') 
        //Profession
        cy.get('input')
        .eq(3).type('Faculty')
        //Name
        cy.get('input')
        .eq(4).type('Gatory')
        //Last name
        cy.get('input')
        .eq(5).type('Marty')
        //Last name
        cy.get('input')
        .eq(6).invoke('removeAttr','type').type('05-02-2011');
    })
    it('Submit registration form', function() {
        cy.get('button').click()       
    })
    it('Login functionality', function() {
        //check this
        cy.visit('http://localhost:4200/')
        })
    it('Login function', function() {
        //Email
        cy.get('input').eq(0).type('user@gatormart.com') 
        //Password
        cy.get('input').eq(1).type('Password1$') 
    })
    it('Submit Button', function() {
        cy.get('button').click()       
    })
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
    //Needs verification
    it('Form text fields', function() {
        //Title
        cy.get('input')
        .eq(2)
        .type('Cypress Electronics') 
        //Seconday Title
        cy.get('input')
        .eq(3)
        .type('Secondary title Electronics') 
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