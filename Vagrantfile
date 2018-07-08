# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    config.vm.box = "Ubuntu1804"
    config.vm.box_url = "https://www.dropbox.com/s/lq0a011b8xhddpg/ubuntu1804lts5212.box?dl=1"
    config.ssh.insert_key = false

    # Host manager setup
    config.hostmanager.enabled				= true
    config.hostmanager.manage_host			= true
    config.hostmanager.manage_guest			= true
    config.hostmanager.ignore_private_ip	= false
    config.hostmanager.include_offline		= true

  	config.vm.define 'JiraStats' do |devbox|
        devbox.vm.network "private_network", ip: "192.168.3.10"
        devbox.vm.hostname = "jirastats.dev"

        devbox.vm.provider "virtualbox" do |vb|
            vb.name = "JiraStats"
        end

        devbox.vm.provision "chef_solo" do |chef|
            chef.cookbooks_path = "./node_modules/rebel-l-sisa/cookbooks"
            chef.roles_path = "./node_modules/rebel-l-sisa/roles"
            chef.environments_path = "./node_modules/rebel-l-sisa/environments"
            chef.data_bags_path = "./node_modules/rebel-l-sisa/data_bags"
            chef.add_role "Default"
            chef.environment = "development"
            chef.add_recipe "Docker"
            chef.add_recipe "GolangCompiler"
            chef.add_recipe "NodeJs"

            chef.json = {
                'Golang' => {
                    'project' => 'github.com/rebel-l/jirastats'
                },
                'System' => {
                    'Iptables' => {
                        'TCP' => {
                            'Ports' => [
                                3000
                            ]
                        }
                    }
                }
            }
        end
    end
end
