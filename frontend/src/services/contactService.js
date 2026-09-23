import request from './api';

export const contactService = {
  /**
   * Submits a message to the backend contact endpoint.
   * @param {Object} payload
   * @param {string} payload.name
   * @param {string} payload.email
   * @param {string} payload.subject
   * @param {string} payload.message
   * @returns {Promise<Object>} Response object
   */
  async submitContact(payload) {
    return await request('/contact', {
      method: 'POST',
      body: JSON.stringify(payload),
    });
  },
};

export default contactService;
