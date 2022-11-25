package existence_conf

import "data-platform-api-invoice-document-creates-rmq-kube/config"

type exconfQueueMapper map[string]string

func NewExconfQueueMapper(c *config.Conf) exconfQueueMapper {
	m := exconfQueueMapper{}
	ecQNames := c.RMQ.QueueToExConf()
	m["BusinessPartnerGeneral"] = ecQNames[1%len(ecQNames)]
	// m["PlantGeneral"] = ecQNames[2%len(ecQNames)]
	return m
}

func (m exconfQueueMapper) getQueueNameByConfContent(content string) string {
	return m[content]
}
