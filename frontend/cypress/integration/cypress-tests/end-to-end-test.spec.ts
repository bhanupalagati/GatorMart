// Objective: Cypress for new software changes - New tests required
/// <reference types ="cypress" />

describe('Complete Functionality', function () {
    it('Sign up', function() {
        cy.visit('http://localhost:4200/signup')
        
    })
    it('Form Fill', function() {
        cy.get('input')
        .eq(1).type('newuser@gatormart.com')
        cy.get('input')
        .eq(2).type('passypass123')
        cy.get('input')
        .eq(3).type('passypass123')
        cy.get('input')
        .eq(4).type('student')
        cy.get('input')
        .eq(5).type('Gator')
        cy.get('input')
        .eq(6).type('Albert')
        cy.get('input')
        .eq(7).invoke('removeAttr','type').type('05-02-2011');
    })
    it('Submit registration form', function() {
        cy.get('button').eq(1).click()       
        })

    it('Sign in', function() {
        cy.visit('http://localhost:4200/')
        
    })
    it('Login function', function() {
        //Email
        cy.get('input')
        .eq(1).type('engineer@user.com') 
        //Password
        cy.get('input')
        .eq(2).type('Test@1234') 
    })
    it('Login Button and remaining functions of site', function() {
        //Login button
        cy.get('button').eq(1).click() 
        cy.wait(3000)  
        //profile button
        cy.get('button').eq(0).click() 
        cy.wait(2000) 
        //Create ad button
        cy.get('button').eq(5).click() 
        cy.wait(2000)
        //images
        //adding the relative path - get the images for the fixtures
        const main_image = 'cypress/fixtures/images/elec.jpg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(0)
        .find('input').selectFile(main_image)  
        const detailed_images = 'cypress/fixtures/images/gadgets.jpeg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(1)
        .find('input').selectFile(detailed_images)  
        // product details
        cy.get('input').eq(3).type('Cypress Electricals') 
        cy.get('input').eq(4).type('Cypress Progress') 
        cy.get('input').eq(6).clear().type('999')
        cy.get('input').eq(7).type('Satellite downlink in progress')
        cy.get('input').eq(8).type('PSLV Satellite deployment')
        cy.wait(1000)  
        cy.get('input').eq(9).type('29.62056457889113')
        cy.get('input').eq(10).type('-82.37617368280584')
        cy.get('input').eq(11).type('Gainesville')
        cy.get('input').eq(12).type('Florida')
        cy.get('mat-select').eq(0).click().get('mat-option').eq(0).click();
        cy.get('input').eq(13).clear().type('1')
        cy.get('mat-select').eq(1).click().get('mat-option').eq(0).click();
        cy.wait(500)
        cy.get('mat-select').eq(2).click().get('mat-option').eq(0).click();
        cy.get('button').eq(2).click()
        //profile button
        cy.get('button').eq(0).click() 
        cy.wait(2000) 
        //my ads button
        cy.get('button').eq(4).click() 
        cy.wait(2000)
        // Click users own products
        cy.get('mat-card').eq(0).click()
        cy.wait(2000)
        //scroll to bottom of page
        cy.scrollTo(0, 1000)
        cy.wait(1500)
        //go to the my ads page
        cy.go('back')
        cy.wait(1000)
        cy.go('back')
        cy.wait(500)
        //Filter
        cy.get('mat-slide-toggle').eq(0).click();
        cy.wait(1500)
        cy.get('input').eq(1).type('Guitar');
        cy.wait(1500)
        cy.get('button').eq(1).click()
        cy.wait(500)
        cy.get('button').eq(2).click()
        cy.wait(1500)
        cy.get('mat-slide-toggle').eq(0).click();
        cy.wait(1500)
        //profile button
        cy.get('button').eq(0).click() 
        cy.wait(1500)
        cy.get('button').eq(6).click() 
    })

}) 