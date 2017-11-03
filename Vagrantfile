# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.box_version = '>= 20160921.0.0'

  config.vm.network "forwarded_port", guest: 8000, host: 8000

  config.vm.synced_folder "./", "/vagrant"

  config.ssh.shell = "bash -c 'BASH_ENV=/etc/profile exec bash'"

  cpus = "1"
  memory = "512" # MB

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--cpus", cpus, "--memory", memory]
    vb.customize ["modifyvm", :id, "--uartmode1", "disconnected"]
    #vb.gui = true
  end

  config.vm.provider "vmware_fusion" do |v, override|
    v.vmx["memsize"] = memory
    v.vmx["numvcpus"] = cpus
  end

  config.vm.provider :cloudstack do |cloudstack, override|
    override.vm.box = "Ubuntu-16.04-SSH-Keys"
    cloudstack.scheme = "https"
    cloudstack.host = "sfucloud.ca"
    cloudstack.path = "/client/api"
    cloudstack.api_key = ENV['CLOUDSTACK_KEY'] || "AAAAAAAAAAAAAAAA"
    cloudstack.secret_key = ENV['CLOUDSTACK_SECRET'] || "AAAAAAAAAAAAAAAA"
    cloudstack.service_offering_name = "sc.t2.micro"
    cloudstack.zone_name = "NML-Zone"
    cloudstack.name = "cmpt470-#{File.basename(Dir.getwd)}-#{Random.new.rand(100)}"
    cloudstack.ssh_user = "ubuntu"
    cloudstack.security_group_names = ['CMPT 470 firewall']
  end

  config.vm.provision "chef_solo" do |chef|
    chef.cookbooks_path = "chef/cookbooks"
    chef.add_recipe "baseconfig"
    chef.channel = "stable"
    chef.version = "12.10.24"
  end
end
