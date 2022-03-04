/// <reference types="cypress" />


//The application should allow a maximum of 4 images for the carousel.
// Test case simultaniously uploads 5 images and 1 for the main image.
// If it fails, the test case is a pass.
describe('Create a new product for carousel test', function () {

    it('Access the page', function() {
        cy.visit('http://localhost:4200/create')
        
    })

    it('Interacting with the 1st image field', function() {
        //adding the relative path - get the images for the fixtures
        const main_image = 'cypress/fixtures/images/elec.jpg'
        cy.get('div [class="ad_form_container__flex__input"]').eq(0)
        .find('input').selectFile(main_image)        
    })

    it('Interacting with the 2nd image field', function() {
        cy.get('div [class="ad_form_container__flex__input"]').eq(1)
        .find('input').selectFile(
            [
                'cypress/fixtures/images/ccmachine.jpg',
                'cypress/fixtures/images/gadgets.jpeg',
                'cypress/fixtures/images/headphones.jpg',
                'cypress/fixtures/images/laptopimage.png',
                'cypress/fixtures/images/ipadimage.jpg'
              ])        
    })
    it('Form text fields', function() {
        //Title
        cy.get('input').eq(2).type('tech Stuff') 
        //Seconday Title
        cy.get('input').eq(3).type('More Electronics') 
        //Price
        cy.get('input').eq(5).clear().type('1550')
        //Simple Description
        cy.get('input').eq(6).type('Cypress image test')
        //Description
        cy.get('input').eq(7).type('Detailed description for cypress images')
        //City
        cy.get('input').eq(8).type('Mysoreville')
        //State
        cy.get('input').eq(9).type('Washington')
        //Latitude
        cy.get('input').eq(10).type('56.1')
        //Longitude
        cy.get('input').eq(11).type('27.1')
        //Condition
        cy.get('input').eq(12).clear().type('New')
        //Age
        cy.get('input').eq(13).clear().type('2')
        //Target
        cy.get('input').eq(15).type('Student')
        //Category
        cy.get('input').eq(16).type('Electronics')
    })
    it('Submit the form', function() {
        cy.get('button').click()       
    })
})

