'use strict';

module.exports = {
    server: {
        port: process.env.PORT || 3000
    },
    
    database: {
        connection: `mysql://${process.env.DB_USER || 'root'}:${process.env.DB_PASSWORD || 'password'}@${process.env.DB_HOST || '127.0.0.1'}:${process.env.DB_PORT || 3306}/${process.env.DB_NAME || 'meu_banco'}`
    },
    
    auth: { 
        secret: process.env.AUTH_SECRET || 'c1c2c3c4c5'
    }
};
