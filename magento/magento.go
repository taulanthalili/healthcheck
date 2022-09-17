package magento

//TEST
func TestMagento(s string) string {
	magentoVersion := MagentoVersion()
	magentoMode := MagentoMode()
	magentoCaches := MagentoCachesStatus()
	magentoVarnish := MagentoVarnishStatus()
	magentoElasticsearch := MagentoElasticsearchStatus()

	magento := string("Magento: " + s + "\n" + magentoVersion + "\n" + magentoMode + "\n" + magentoCaches + "\n" + magentoVarnish + "\n" + magentoElasticsearch)
	return string(magento)
}

//Check the Magento Version
func MagentoVersion() string {
	return string("Magento Version")
}

//Check if Magento mode is "default", "production"...
func MagentoMode() string {
	return string("Magento Mode")
}

//Check if the Caches are enabled
func MagentoCachesStatus() string {
	return string("Magento Caches Status")
}

//Check if the Varnish Magento Module is enabled
func MagentoVarnishStatus() string {
	return string("Magento Varnish Status")
}

//Check if the Elasticsearch Magento Module is enabled
func MagentoElasticsearchStatus() string {
	return string("Magento Elasticsearch status")
}